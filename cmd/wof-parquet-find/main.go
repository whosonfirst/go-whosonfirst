package main

import (
	"log"
	"flag"
	"encoding/json"
	"os"
	
	"github.com/whosonfirst/go-whosonfirst/v4/parquet"
)

func main() {

	var parquet_uri string
	var id int64

	flag.StringVar(&parquet_uri, "parquet-uri", "", "")
	flag.Int64Var(&id, "id", 0, "")

	flag.Parse()

	rec, err := parquet.FindRecord(parquet_uri, id)

	if err != nil {
		log.Fatal(err)
	}

	f, err := rec.AsGeoJSON()

	if err != nil {
		log.Fatal(err)
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(f)

	if err != nil {
		log.Fatal(err)
	}
	
}
