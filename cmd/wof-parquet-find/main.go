package main

// > ./bin/wof-parquet-find -parquet-uri /usr/local/data/whosonfirst-parquet/whosonfirst-data-admin-us.parquet -id 1477727311 | jq '.properties["wof:name"]'
// "Appalachian National Scenic Trail"

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/whosonfirst/go-whosonfirst/v4/parquet"
)

func main() {

	var parquet_uri string
	var id int64

	flag.StringVar(&parquet_uri, "parquet-uri", "The URI of the Parquet file to search for a record in.", "")
	flag.Int64Var(&id, "id", -1, "The Who's On First ID to search for.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Finds a Who's On First record by ID in a Parquet file.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	rec, err := parquet.FindRecord(parquet_uri, id)

	if err != nil {
		log.Fatalf("Failed to locate record, %v", err)
	}

	f, err := rec.AsGeoJSON()

	if err != nil {
		log.Fatalf("Failed to cast record as GeoJSON, %v", err)
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(f)

	if err != nil {
		log.Fatalf("Failed to JSON-encode record, %w", err)
	}

}
