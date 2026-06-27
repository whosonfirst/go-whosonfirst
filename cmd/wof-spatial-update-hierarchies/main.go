package main

import (
	"context"
	"log"

	"github.com/whosonfirst/go-whosonfirst/v4/app/spatial/hierarchy/update"
)

func main() {

	ctx := context.Background()
	err := update.Run(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
