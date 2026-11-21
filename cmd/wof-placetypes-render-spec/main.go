package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/whosonfirst/go-whosonfirst/v4/placetypes"
	"github.com/whosonfirst/go-whosonfirst/v4/placetypes/draw"
)

func main() {

	path := flag.String("path", "placetypes.png", "...")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Renders the placetype specification as a graph.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	spec, err := placetypes.DefaultWOFPlacetypeSpecification()

	if err != nil {
		log.Fatalf("Failed to load specification, %v", err)
	}

	err = draw.DrawPlacetypesGraphToFile(spec, *path)

	if err != nil {
		log.Fatalf("Failed to draw placetypes graph, %v", err)
	}
}
