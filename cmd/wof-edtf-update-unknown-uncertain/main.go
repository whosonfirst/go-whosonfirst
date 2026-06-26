package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	"github.com/tidwall/gjson"
	"github.com/whosonfirst/go-whosonfirst/v4/edtf"
	"github.com/whosonfirst/go-whosonfirst/v4/iterate"
	wof_writer "github.com/whosonfirst/go-whosonfirst/v4/writer"
	"github.com/whosonfirst/go-writer/v3"
)

func main() {

	emitter_schemes := ""
	iterator_desc := fmt.Sprintf("A valid whosonfirst/go-whosonfirst-iterate/v2 URI. Supported emitter URI schemes are: %s", emitter_schemes)

	iterator_uri := flag.String("iterator-uri", "repo://", iterator_desc)

	writer_uri := flag.String("writer-uri", "null://", "A valid whosonfirst/go-writer URI.")

	flag.Parse()

	uris := flag.Args()

	ctx := context.Background()

	wr, err := writer.NewWriter(ctx, *writer_uri)

	if err != nil {
		log.Fatalf("Failed to create writer for '%s', %v", *writer_uri, err)
	}

	iter, err := iterate.NewIterator(ctx, *iterator_uri)

	if err != nil {
		log.Fatal(err)
	}

	for rec, err := range iter.Iterate(ctx, uris...) {

		if err != nil {
			log.Fatal(err)
		}

		defer rec.Body.Close()

		body, err := io.ReadAll(rec.Body)

		if err != nil {
			log.Fatal(err)
		}

		id_rsp := gjson.GetBytes(body, "properties.wof:id")

		if !id_rsp.Exists() {
			log.Fatalf("missing wof:id property")
		}

		id := id_rsp.Int()

		changed, body, err := edtf.UpdateBytes(body)

		if err != nil {
			log.Fatalf("Failed to apply EDTF updates to %d, %v", id, err)
		}

		if !changed {
			continue
		}
		_, err = wof_writer.WriteBytes(ctx, wr, body)

		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Updated %d (%s)\n", id, rec.Path)
	}
}
