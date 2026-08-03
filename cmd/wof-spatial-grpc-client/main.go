package main

import (
	"context"
	"log"

	_ "github.com/whosonfirst/go-whosonfirst/v4/spatial/pmtiles"
	_ "github.com/whosonfirst/go-whosonfirst/v4/spatial/sqlite"	
	
	"github.com/whosonfirst/go-whosonfirst/v4/app/spatial/grpc/client"
)

func main() {

	ctx := context.Background()
	err := client.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to run client, %v", err)
	}
}
