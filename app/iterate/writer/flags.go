package writer

import (
	"flag"
	"fmt"
	"os"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/sfomuseum/go-flags/multi"
)

var writer_uris multi.MultiCSVString
var iterator_uri string
var monitor_uri string
var forgiving bool
var verbose bool

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("es")

	fs.Var(&writer_uris, "writer-uri", "One or more valid whosonfirst/go-writer/v4 URIs, each encoded as a gocloud.dev/runtimevar URI.")
	fs.StringVar(&iterator_uri, "iterator-uri", "repo://", "A valid whosonfirst/go-whosonfirst-iterate/v4 URI.")
	fs.StringVar(&monitor_uri, "monitor-uri", "counter://PT60S", "A valid sfomuseum/go-timings URI.")
	fs.BoolVar(&forgiving, "forgiving", false, "Be \"forgiving\" of failed writes, logging the issue(s) but not triggering errors")
	fs.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Sends iterated records to a writer.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options] path(N) path(N)\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n\n")
		fs.PrintDefaults()
	}

	return fs
}
