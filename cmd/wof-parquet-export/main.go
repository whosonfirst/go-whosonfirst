package main

import (
	"context"
	"log"
	"fmt"
	"log/slog"
	
	_ "github.com/whosonfirst/go-whosonfirst/v4/iterate/git"
	
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/whosonfirst/go-whosonfirst-github/organizations"	
	"github.com/whosonfirst/go-whosonfirst/v4/iterate"
	"github.com/whosonfirst/go-whosonfirst/v4/parquet"
)

func main() {

	var iterator_uri string

	fs := flagset.NewFlagSet("parquet")
	fs.StringVar(&iterator_uri, "iterator-uri", "git:///tmp", "A registered whosonfirst/go-whosonfirst/v4/iterate.Iterator URI.")

	flagset.Parse(fs)

	ctx := context.Background()

	list_opts := &organizations.ListOptions{
		Prefix: []string{
			"whosonfirst-data-admin",
		},
		Exclude: []string{
			"whosonfirst-data-venue",
		},
		ExcludeArchived: true,
	}

	repos, err := organizations.ListRepos("whosonfirst-data", list_opts)

	if err != nil {
		log.Fatalf("Failed to list repos, %v", err)
	}
	
	for _, repo := range repos {
	
		writer_uri := fmt.Sprintf("/usr/local/data/whosonfirst-parquet/%s.parquet", repo)
		iterator_source := fmt.Sprintf("https://github.com/whosonfirst-data/%s.git", repo)

		logger := slog.Default()
		logger = logger.With("source", iterator_source)
		logger = logger.With("target", writer_uri)

		logger.Info("Start export")
		
		wr, err := parquet.NewWriter(ctx, writer_uri)
		
		if err != nil {
			log.Fatalf("Failed to create new writer, %v", err)
		}
		
		iter, err := iterate.NewIterator(ctx, iterator_uri)
		
		if err != nil {
			log.Fatalf("Failed to create new iterator, %v", err)
		}

		for rec, err := range iter.Iterate(ctx, iterator_source) {

			if err != nil {
				log.Fatalf("Iterator yielded an error, %v", err)
			}
			
			_, err = wr.WriteFromReader(rec.Body)
			
			rec.Body.Close()
			
			if err != nil {
				log.Fatalf("Failed to write row %s, %v", rec.Path, err)
			}
		}
		
		err = wr.Close()
		
		if err != nil {
			log.Fatalf("Failed to close after writing, %v", err)
		}

		logger.Info("Completed export")		
	}
}
