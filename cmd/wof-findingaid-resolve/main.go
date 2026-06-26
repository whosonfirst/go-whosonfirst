package main

import (
	"context"
	"log"

	"github.com/whosonfirst/go-whosonfirst/v4/app/findingaid/resolve"
)

func main() {

	ctx := context.Background()
	err := resolve.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to run resolve tool, %v", err)
	}
}
