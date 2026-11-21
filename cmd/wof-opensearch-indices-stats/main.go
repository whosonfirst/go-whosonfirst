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

	fs := flagset.NewFlagSet("opensearch")

	client.AppendClientFlags(fs)

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Retrieves OpenSearch index statistics.\n")
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

	req := &opensearchapi.IndicesStatsReq{
		Params: opensearchapi.IndicesStatsParams{
			Pretty: true,
		},
	}

	rsp, err := os_client.Indices.Stats(ctx, req)

	if err != nil {
		log.Fatalf("Failed to execute request, %v", err)
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(rsp)

	if err != nil {
		log.Fatalf("Failed to copy response, %v", err)
	}

	os.Exit(0)
}
