package id

import (
	"context"
	"fmt"
	"testing"
)

func TestNewID(t *testing.T) {

	_, err := NewID()

	if err != nil {
		t.Fatalf("Failed to create new ID, %v", err)
	}
}

func TestNextInt(t *testing.T) {

	ctx := context.Background()
	pr, err := NewProvider(ctx)

	if err != nil {
		t.Fatalf("Failed to create provider, %v", err)
	}

	id, err := pr.NewID(ctx)

	if err != nil {
		t.Fatalf("Failed to generate new ID, %v", err)
	}

	fmt.Println(id)
}

func TestNextIntWithURI(t *testing.T) {

	ctx := context.Background()

	uri := "proxy://?provider=whosonfirst://&minimum=10"
	pr, err := NewProviderWithURI(ctx, uri)

	if err != nil {
		t.Fatalf("Failed to create provider for %s, %v", uri, err)
	}

	id, err := pr.NewID(ctx)

	if err != nil {
		t.Fatalf("Failed to generate new ID, %v", err)
	}

	fmt.Println(id)
}
