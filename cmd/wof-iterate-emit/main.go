package main

import (
	"context"
	"log"

	"github.com/whosonfirst/go-whosonfirst/v4/app/iterate/emit"
)

func main() {

	ctx := context.Background()
	err := emit.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to emit records, %v", err)
	}
}
