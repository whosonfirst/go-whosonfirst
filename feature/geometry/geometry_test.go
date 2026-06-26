package geometry

import (
	"io"
	"strings"
	"testing"
)

var valid_geom = strings.NewReader(`{"geometry": { "type": "Point", "coordinates": [ 0.0, 0.0 ] }}`)

var missing_geom = strings.NewReader(`{"properties":{ }}`)

func TestValidGeometry(t *testing.T) {

	body, err := io.ReadAll(valid_geom)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	geom, err := Geometry(body)

	if err != nil {
		t.Fatalf("Expect data (valid) failed")
	}

	orb_geom := geom.Geometry()

	if orb_geom.GeoJSONType() != "Point" {
		t.Fatal("Unexpected geometry type")
	}
}

func TestMissingGeometry(t *testing.T) {

	body, err := io.ReadAll(missing_geom)

	if err != nil {
		t.Fatalf("Failed to read data (missing), %v", err)
	}

	_, err = Geometry(body)

	if err == nil {
		t.Fatalf("Expect data (missing) to fail")
	}
}
