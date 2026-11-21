package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
	"github.com/sfomuseum/go-flags/flagset"
	"github.com/whosonfirst/go-whosonfirst/v4/database/opensearch/client"
)

func main() {

	var os_index string
	var settings string

	fs := flagset.NewFlagSet("opensearch")

	fs.StringVar(&os_index, "opensearch-index", "", "...")
	fs.StringVar(&settings, "settings", "", "...")

	client.AppendClientFlags(fs)

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Creates an OpenSearch index.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n\n")
		fs.PrintDefaults()
	}

	flagset.Parse(fs)

	ctx := context.Background()

	os_client, err := client.NewClientFromFlagSet(ctx, fs)

	if err != nil {
		log.Fatalf("Failed to create Opensearch client, %v", err)
	}

	req := opensearchapi.IndicesCreateReq{
		Index: os_index,
		Params: opensearchapi.IndicesCreateParams{
			Pretty: true,
		},
	}

	if settings != "" {

		r, err := os.Open(settings)

		if err != nil {
			log.Fatalf("Failed to open settings for reading, %v", err)
		}

		defer r.Close()
		req.Body = r
	}

	rsp, err := os_client.Indices.Create(ctx, req)

	if err != nil {
		log.Fatalf("Failed to create index '%s', %v", os_index, err)
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(rsp)

	if err != nil {
		log.Fatalf("Failed to copy response, %v", err)
	}

	os.Exit(0)
}
