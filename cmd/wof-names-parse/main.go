package main

import (
	"flag"
	"log"

	"github.com/whosonfirst/go-whosonfirst/v4/names/tags"
	"github.com/whosonfirst/go-whosonfirst/v4/names/utils"
)

func main() {

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
