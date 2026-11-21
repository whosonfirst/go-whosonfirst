package export

import (
	"flag"
	"fmt"
	"os"

	"github.com/sfomuseum/go-flags/flagset"
)

var exporter_uri string

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("export")

	fs.StringVar(&exporter_uri, "exporter-uri", "whosonfirst://", "A valid whosonfirst/go-whosonfirst-export/v3.Exporter URI")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Exports Who's On First records.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options] path(N) path(N)\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n\n")
		fs.PrintDefaults()
	}

	return fs
}
