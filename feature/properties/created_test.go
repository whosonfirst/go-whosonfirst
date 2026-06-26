package properties

import (
	"io"
	"strings"
	"testing"
)

var valid_created = strings.NewReader(`{"properties":{"wof:created": 1635460530 }}`)

var missing_created = strings.NewReader(`{"properties":{ }}`)

func TestValidCreated(t *testing.T) {

	body, err := io.ReadAll(valid_created)

	if err != nil {
		t.Fatalf("Failed to read data (valid), %v", err)
	}

	created := Created(body)

	if err != nil {
		t.Fatalf("Expect data (valid) failed")
	}

	if created != 1635460530 {
		t.Fatal("Invalid CREATED (valid)")
	}

}

func TestMissingCreated(t *testing.T) {

	body, err := io.ReadAll(missing_created)

	if err != nil {
		t.Fatalf("Failed to read data (missing), %v", err)
	}

	created := Created(body)

	if err != nil {
		t.Fatalf("Expect data (missing) failed")
	}

	if created != -1 {
		t.Fatal("Invalid CREATED (missing)")
	}
}
