package main

import (
	"context"
	"log"

	_ "github.com/whosonfirst/go-whosonfirst/v4/database/sql"
	_ "github.com/whosonfirst/go-whosonfirst/v4/database/sql/writer"

	"github.com/whosonfirst/go-whosonfirst-iterwriter/v4/app/iterwriter"
)

func main() {
	ctx := context.Background()
	err := iterwriter.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to run iterwriter, %v", err)
	}
}
