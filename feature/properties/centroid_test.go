package properties

import (
	"io"
	"strings"
	"testing"
)

var lbl_centroid = strings.NewReader(`{"properties": {"lbl:latitude":41.933211, "lbl:longitude":-87.677286, "reversegeo:latitude":41.943403, "reversegeo:longitude":-87.665224}}`)

var reversegeo_centroid = strings.NewReader(`{"properties": {"reversegeo:latitude":41.943403, "reversegeo:longitude":-87.665224, "geom:latitude":41.943648, "geom:longitude":-87.665224}}`)

var geom_centroid = strings.NewReader(`{"properties": {"geom:latitude":41.943648, "geom:longitude":-87.665224}}`)

var null_centroid = strings.NewReader(`{"properties": {}}`)

func TestLblCentroid(t *testing.T) {

	body, err := io.ReadAll(lbl_centroid)

	if err != nil {
		t.Fatalf("Failed to read data , %v", err)
	}

	pt, source, err := Centroid(body)

	if err != nil {
		t.Fatalf("Failed to derive centroid (lbl), %v", err)
	}

	if source != "lbl" {
		t.Fatalf("Unexpected source")
	}

	if pt.Lat() != 41.933211 {
		t.Fatalf("Unexpected latitude")
	}

	if pt.Lon() != -87.677286 {
		t.Fatalf("Unexpected longitude")
	}

}

func TestReversegeoCentroid(t *testing.T) {

	body, err := io.ReadAll(reversegeo_centroid)

	if err != nil {
		t.Fatalf("Failed to read data , %v", err)
	}

	pt, source, err := Centroid(body)

	if err != nil {
		t.Fatalf("Failed to derive centroid, %v", err)
	}

	if source != "reversegeo" {
		t.Fatalf("Unexpected source")
	}

	if pt.Lat() != 41.943403 {
		t.Fatalf("Unexpected latitude")
	}

	if pt.Lon() != -87.665224 {
		t.Fatalf("Unexpected longitude")
	}

}

func TestGeomCentroid(t *testing.T) {

	body, err := io.ReadAll(geom_centroid)

	if err != nil {
		t.Fatalf("Failed to read data , %v", err)
	}

	pt, source, err := Centroid(body)

	if err != nil {
		t.Fatalf("Failed to derive centroid, %v", err)
	}

	if source != "geom" {
		t.Fatalf("Unexpected source")
	}

	if pt.Lat() != 41.943648 {
		t.Fatalf("Unexpected latitude")
	}

	if pt.Lon() != -87.665224 {
		t.Fatalf("Unexpected longitude")
	}

}

func TestNullCentroid(t *testing.T) {

	body, err := io.ReadAll(null_centroid)

	if err != nil {
		t.Fatalf("Failed to read data , %v", err)
	}

	pt, source, err := Centroid(body)

	if err != nil {
		t.Fatalf("Failed to derive centroid (null), %v", err)
	}

	if source != "nullisland" {
		t.Fatalf("Unexpected source")
	}

	if pt.Lat() != 0.0 {
		t.Fatalf("Unexpected latitude")
	}

	if pt.Lon() != 0.0 {
		t.Fatalf("Unexpected longitude")
	}

}
