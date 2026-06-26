package main

import (
	"flag"
	"log"

	"github.com/sfomuseum/go-flags/multi"
	"github.com/whosonfirst/go-whosonfirst/v4/placetypes"
)

func main() {

	var roles multi.MultiString
	flag.Var(&roles, "role", "...")

	flag.Parse()

	for _, str_pt := range flag.Args() {

		pt, err := placetypes.GetPlacetypeByName(str_pt)

		if err != nil {
			log.Fatal(err)
		}

		children := placetypes.Children(pt)

		for i, p := range children {
			log.Println(i, p.Name)
		}
	}
}
