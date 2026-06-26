package main

import (
	"context"
	"log"

	"github.com/whosonfirst/go-whosonfirst/v4/app/derivatives/server"
)

func main() {

	ctx := context.Background()
	err := server.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to run server, %v", err)
	}
}
