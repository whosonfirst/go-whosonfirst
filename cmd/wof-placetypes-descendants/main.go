package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sfomuseum/go-flags/multi"
	"github.com/whosonfirst/go-whosonfirst/v4/placetypes"
)

func main() {

	var roles multi.MultiString
	flag.Var(&roles, "role", "...")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Lists descendant placetypes, with optional role filtering.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options] path(N) path(N)\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	for _, str_pt := range flag.Args() {

		pt, err := placetypes.GetPlacetypeByName(str_pt)

		if err != nil {
			log.Fatal(err)
		}

		var descendants []*placetypes.WOFPlacetype

		if len(roles) > 0 {
			descendants = placetypes.DescendantsForRoles(pt, roles)
		} else {
			descendants = placetypes.Descendants(pt)
		}

		for i, p := range descendants {
			log.Println(i, p.Name)
		}
	}
}
