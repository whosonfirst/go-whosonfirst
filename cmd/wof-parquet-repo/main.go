package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/whosonfirst/go-whosonfirst/v4/findingaid/parquet"
)

func main() {

	var parquet_uri string
	var id int64

	flag.StringVar(&parquet_uri, "parquet-uri", "The URI of the Parquet file to search for a record in.", "")
	flag.Int64Var(&id, "id", -1, "The Who's On First URI (ID) to search for.")

	flag.Parse()

	repo, err := parquet.GetRepo(parquet_uri, id)

	if err != nil {
		log.Fatalf("Failed to locate record, %v", err)
	}

	fmt.Println(repo)
}
