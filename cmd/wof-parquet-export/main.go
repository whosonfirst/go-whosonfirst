package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"

	_ "github.com/whosonfirst/go-whosonfirst/v4/iterate/git"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
	"github.com/whosonfirst/go-whosonfirst-github/organizations"
	"github.com/whosonfirst/go-whosonfirst/v4/iterate"
	"github.com/whosonfirst/go-whosonfirst/v4/parquet"
)

func main() {

	var iterator_uri string
	var writer_root string
	var org string
	var repos multi.MultiString
	var prefixes multi.MultiString
	var exclude multi.MultiString
	var verbose bool

	fs := flagset.NewFlagSet("parquet")
	fs.StringVar(&iterator_uri, "iterator-uri", "git:///tmp", "A registered whosonfirst/go-whosonfirst/v4/iterate.Iterator URI.")
	fs.StringVar(&writer_root, "writer-root", "/usr/local/data/whosonfirst-parquet", "The path on the local filesystem where Parquet files should be written to.")
	fs.StringVar(&org, "org", "whosonfirst-data", "The organization to use when cloning repos.")
	fs.Var(&prefixes, "prefix", "Zero or more repository prefixe to filter by (if -repo flag is empty). Default is [whosonfirst-data-admin].")
	fs.Var(&exclude, "exclude", "Zero or more repository prefixe to exclude by (if -repo flag is empty). Default is [].")
	fs.Var(&repos, "repo", "Zero or more repos to clone and create Parquet exports for. If empty then repos will be limited to those starting in whosonfirst-data-admin.")
	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Write data from one or more Who's On First style repositories to Parquet files.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t%s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n")
		fs.PrintDefaults()
	}

	flagset.Parse(fs)

	if verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	ctx := context.Background()

	abs_root, err := filepath.Abs(writer_root)

	if err != nil {
		log.Fatalf("Failed to derive absolute path for writer root, %v", err)
	}

	if len(repos) == 0 {

		if len(prefixes) == 0 {
			prefixes = []string{
				"whosonfirst-data-admin",
			}
		}

		list_opts := &organizations.ListOptions{
			Prefix:          prefixes,
			Exclude:         exclude,
			ExcludeArchived: true,
		}

		list_repos, err := organizations.ListRepos(org, list_opts)

		if err != nil {
			log.Fatalf("Failed to list repos, %v", err)
		}

		repos = list_repos
	}

	for _, repo := range repos {

		writer_fname := fmt.Sprintf("%s.parquet", repo)
		writer_uri := filepath.Join(abs_root, writer_fname)

		iterator_source := fmt.Sprintf("https://github.com/%s/%s.git", org, repo)

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
