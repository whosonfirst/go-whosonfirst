package main

import (
	"context"
	"log"

	iterwriter "github.com/whosonfirst/go-whosonfirst/v4/app/iterate/writer"
)

func main() {
	ctx := context.Background()
	err := iterwriter.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to run iterwriter, %v", err)
	}
}
