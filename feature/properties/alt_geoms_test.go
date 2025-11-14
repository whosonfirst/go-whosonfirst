package properties

import (
	"io"
	"strings"
	"testing"
)

var valid_geoms = strings.NewReader(`{"properties":{"src:geom_alt": ["quattroshapes","debug"] }}`)

var missing_geoms = strings.NewReader(`{"properties":{ }}`)

func TestValidAltGeoms(t *testing.T) {

	body, err := io.ReadAll(valid_geoms)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	alt_geoms, err := AltGeometries(body)

	if err != nil {
		t.Fatalf("Expect data (valid) failed")
	}

	if alt_geoms[0] != "quattroshapes" {
		t.Fatal("Invalid alt geometry (valid)")
	}

}

func TestMissingGeoms(t *testing.T) {

	body, err := io.ReadAll(missing_geoms)

	if err != nil {
		t.Fatalf("Failed to read data (missing), %v", err)
	}

	alt_geoms, err := AltGeometries(body)

	if err != nil {
		t.Fatalf("Expect data (valid) failed")
	}

	if len(alt_geoms) != 0 {
		t.Fatalf("Expect data (missing) to fail")
	}
}
