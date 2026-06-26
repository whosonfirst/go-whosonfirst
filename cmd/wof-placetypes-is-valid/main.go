package main

import (
	"flag"
	"log"

	"github.com/whosonfirst/go-whosonfirst/v4/placetypes"
)

func main() {

	flag.Parse()

	for _, pt := range flag.Args() {
		log.Printf("%s\t%t\n", pt, placetypes.IsValidPlacetype(pt))
	}
}
