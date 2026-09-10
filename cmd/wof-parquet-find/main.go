package main

// go run cmd/wof-parquet-find/main.go -parquet-uri /usr/local/data/whosonfirst-parquet/whosonfirst-data-admin-us.parquet -uri 85922583 | show -

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/whosonfirst/go-whosonfirst/v4/parquet"
	wof_uri "github.com/whosonfirst/go-whosonfirst/v4/uri"
)

func main() {

	var parquet_uri string
	var uri string

	flag.StringVar(&parquet_uri, "parquet-uri", "The URI of the Parquet file to search for a record in.", "")
	flag.StringVar(&uri, "uri", "", "The Who's On First URI (ID) to search for.")

	flag.Parse()

	id, uri_args, err := wof_uri.ParseURI(uri)

	if err != nil {
		log.Fatalf("Failed to parse URI, %v", err)
	}

	record, err := parquet.FindRecord(parquet_uri, id, uri_args)

	if err != nil {
		log.Fatalf("Failed to locate record, %v", err)
	}

	fc, err := parquet.RecordsAsFeatureCollection(record)

	if err != nil {
		log.Fatalf("Failed to cast record as GeoJSON, %v", err)
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(fc)

	if err != nil {
		log.Fatalf("Failed to JSON-encode record, %w", err)
	}

}
