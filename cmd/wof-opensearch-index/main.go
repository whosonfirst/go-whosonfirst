package main

import (
	"context"
	"log"

	// _ "github.com/whosonfirst/go-whosonfirst-iterate-git/v3"
	_ "github.com/whosonfirst/go-whosonfirst/v4/database/opensearch/writer"

	iterwriter "github.com/whosonfirst/go-whosonfirst/v4/app/iterate/writer"
)

func main() {

	ctx := context.Background()
	err := iterwriter.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to iterate, %v", err)
	}

}
