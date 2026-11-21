package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/whosonfirst/go-whosonfirst/v4/iterate"
	"github.com/whosonfirst/go-whosonfirst/v4/parquet"
)

func main() {

	var writer_uri string
	var iterator_uri string

	fs := flagset.NewFlagSet("parquet")
	fs.StringVar(&writer_uri, "writer-uri", "", "The path to a file where Parquet data should be written. If \"-\" then data will be written to STDOUT.")
	fs.StringVar(&iterator_uri, "iterator-uri", "repo://", "A registered whosonfirst/go-whosonfirst/v4/iterate.Iterator URI.")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Converts records from a configured iterator to Parquet rows and writes them to a configured destination.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options] path(N) path(N)\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n\n")
		fs.PrintDefaults()
	}

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

		_, err = parquet.WriteFromReader(wr, rec.Body)

		rec.Body.Close()

		if err != nil {
			log.Fatalf("Failed to write row %s, %v", rec.Path, err)
		}
	}

	err = wr.Close()

	if err != nil {
		log.Fatalf("Failed to close after writing, %v", err)
	}
}
