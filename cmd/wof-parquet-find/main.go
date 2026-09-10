package main

// go run cmd/wof-parquet-find/main.go -parquet-uri /usr/local/data/whosonfirst-parquet/whosonfirst-data-admin-us.parquet -id 85922583 | show -

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/whosonfirst/go-whosonfirst/v4/parquet"
)

func main() {

	var parquet_uri string
	var id int64

	flag.StringVar(&parquet_uri, "parquet-uri", "The URI of the Parquet file to search for a record in.", "")
	flag.Int64Var(&id, "id", -1, "The Who's On First ID to search for.")

	flag.Parse()

	records, err := parquet.FindRecords(parquet_uri, id)

	if err != nil {
		log.Fatalf("Failed to locate record, %v", err)
	}

	fc, err := parquet.RecordsAsFeatureCollection(records)

	if err != nil {
		log.Fatalf("Failed to cast record as GeoJSON, %v", err)
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(fc)

	if err != nil {
		log.Fatalf("Failed to JSON-encode record, %w", err)
	}

}
