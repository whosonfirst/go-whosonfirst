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
		fmt.Fprintf(os.Stderr, "Lists known Who's On First placetypes.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	pt, err := placetypes.Placetypes()

	if err != nil {
		log.Fatal(err)
	}

	for i, p := range pt {
		log.Println(i, p.Name)
	}
}
