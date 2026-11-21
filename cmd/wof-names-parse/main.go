package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/whosonfirst/go-whosonfirst/v4/names/tags"
	"github.com/whosonfirst/go-whosonfirst/v4/names/utils"
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Parses Who's On First name language tags.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options] path(N) path(N)\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	for _, raw := range flag.Args() {

		log.Println(raw)
		langtag, err := tags.NewLangTag(raw)

		if err != nil {
			log.Fatal(err)
		}

		log.Println(langtag.String())
		log.Println(utils.ToRFC5646(langtag.String()))
	}
}
