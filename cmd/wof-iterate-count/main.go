package main

import (
	"context"
	"log"

	_ "github.com/whosonfirst/go-whosonfirst/v4/iterate/git"
	_ "github.com/whosonfirst/go-whosonfirst/v4/iterate/parquet"

	"github.com/whosonfirst/go-whosonfirst/v4/app/iterate/count"
)

func main() {

	ctx := context.Background()
	err := count.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to count records, %v", err)
	}
}
