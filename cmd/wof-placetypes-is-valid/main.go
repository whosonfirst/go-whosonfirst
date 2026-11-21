package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/whosonfirst/go-whosonfirst/v4/placetypes"
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Checks whether a name identifies a known Who's On First placetype.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options] path(N) path(N)\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	for _, pt := range flag.Args() {
		log.Printf("%s\t%t\n", pt, placetypes.IsValidPlacetype(pt))
	}
}
