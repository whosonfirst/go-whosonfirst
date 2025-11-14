package feature

import (
	"errors"
	"testing"
)

func TestPropertyNotFoundError(t *testing.T) {

	err := PropertyNotFoundError("wof:id")

	if !IsPropertyNotFoundError(err) {
		t.Fatalf("Not a property not found error")
	}

	if err.Property() != "wof:id" {
		t.Fatalf("Invalid property: %s", err.Property())
	}

	err2 := errors.New("Some other error")

	if IsPropertyNotFoundError(err2) {
		t.Fatalf("Should not be a property not found error")
	}

}
