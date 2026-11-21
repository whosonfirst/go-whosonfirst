package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/whosonfirst/go-whosonfirst/v4/findingaid/provider"
)

func main() {

	provider_uri := flag.String("provider-uri", "github://whosonfirst-data", "...")

	uri_template := flag.String("uri-template", "", "...")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Lists iterator sources supplied by a finding aid provider.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	ctx := context.Background()

	pr, err := provider.NewProvider(ctx, *provider_uri)

	if err != nil {
		log.Fatalf("Failed to create new provider, %v", err)
	}

	var sources []string

	if *uri_template != "" {
		sources, err = pr.IteratorSourcesWithURITemplate(ctx, *uri_template)
	} else {
		sources, err = pr.IteratorSources(ctx)
	}

	if err != nil {
		log.Fatalf("Failed to derive sources, %v", err)
	}

	for _, s := range sources {
		fmt.Println(s)
	}
}
