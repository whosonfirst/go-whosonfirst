package main

import (
	"context"
	"log"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/whosonfirst/go-whosonfirst/v4/iterate"
	"github.com/whosonfirst/go-whosonfirst/v4/parquet"
)

func main() {

	var writer_uri string
	var iterator_uri string

	fs := flagset.NewFlagSet("parquet")
	fs.StringVar(&writer_uri, "writer-uri", "", "...")
	fs.StringVar(&iterator_uri, "iterator-uri", "repo://", "...")

	flagset.Parse(fs)

	ctx := context.Background()

	wr, err := parquet.NewWriter(ctx, writer_uri)

	if err != nil {
		log.Fatalf("Failed to create new writer, %v", err)
	}

	iter, err := iterate.NewIterator(ctx, iterator_uri)

	if err != nil {
		log.Fatalf("Failed to create new iterator, %v", err)
	}

	iterator_sources := fs.Args()

	for rec, err := range iter.Iterate(ctx, iterator_sources...) {

		if err != nil {
			log.Fatalf("Iterator yielded an error, %v", err)
		}

		_, err = wr.WriteFromReader(rec.Body)

		if err != nil {
			log.Fatalf("Failed to write row %s, %v", rec.Path, err)
		}
	}

	err = wr.Close()

	if err != nil {
		log.Fatalf("Failed to close after writing, %v", err)
	}
}
