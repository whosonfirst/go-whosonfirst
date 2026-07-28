package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/paulmach/orb/geojson"
	"github.com/whosonfirst/go-whosonfirst/v4/spatial/geo"
)

func main() {

	flag.Parse()

	for _, path := range flag.Args() {

		body, err := os.ReadFile(path)

		if err != nil {
			log.Fatalf("Failed to read %s, %v", path, err)
		}

		f, err := geojson.UnmarshalFeature(body)

		if err != nil {
			log.Fatalf("Failed to unmarsal %s, %v", path, err)
		}

		pt, ok := geo.FindInnerPoint(f.Geometry)

		if !ok {
			log.Fatalf("Failed to derive inner point for %s", path)
		}

		pt_f := geojson.NewFeature(pt)
		pt_f.Properties = map[string]any{
			"hello": "world",
		}

		enc := json.NewEncoder(os.Stdout)
		enc.Encode(pt_f)
	}

}
