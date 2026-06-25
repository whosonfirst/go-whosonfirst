// fetch is a command line tool to retrieve one or more Who's on First records and, optionally, their ancestors.
package main

import (
	"context"
	"log"

	_ "github.com/whosonfirst/go-whosonfirst/v4/findingaid/reader"

	"github.com/whosonfirst/go-whosonfirst/v4/app/fetch/records"
)

func main() {

	ctx := context.Background()
	err := records.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to fetch records, %v", err)
	}
}
