package main

import (
	"context"
	"log"

	"github.com/whosonfirst/go-whosonfirst/v4/app/spatial/pip"
)

func main() {

	ctx := context.Background()
	err := pip.Run(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
