// Package export prepares Who's On First records for storage by applying
// defaults, validating properties, and formatting the result.
//
// # Example
//
//	package main
//
//	import (
//		"context"
//		"log"
//		"os"
//
//		"github.com/whosonfirst/go-whosonfirst/v4/export"
//	)
//
//	func main() {
//		ctx := context.Background()
//		body, err := os.ReadFile("some.geojson")
//		if err != nil {
//			log.Fatal(err)
//		}
//		if _, err := export.WriteExportIfChanged(ctx, body, os.Stdout); err != nil {
//			log.Fatal(err)
//		}
//	}
package export
