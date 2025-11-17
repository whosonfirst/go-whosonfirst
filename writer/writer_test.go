package writer

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/paulmach/orb/geojson"
	go_writer "github.com/whosonfirst/go-writer/v3"	
)

func TestWriteFeature(t *testing.T) {

	ctx := context.Background()

	body, err := load_feature(ctx)

	if err != nil {
		t.Fatalf("Failed to load feature, %v", err)
	}

	f, err := geojson.UnmarshalFeature(body)

	if err != nil {
		t.Fatalf("Failed to unmarshal feature, %v", err)
	}

	wr, err := go_writer.NewWriter(ctx, "null://")

	if err != nil {
		t.Fatalf("Failed to create new writer, %v", err)
	}

	id, err := WriteFeature(ctx, wr, f)

	if err != nil {
		t.Fatalf("Failed to write feature, %v", err)
	}

	if id != 101736545 {
		t.Fatalf("Unexpected ID returned: %d", id)
	}
}

func TestWriteBytes(t *testing.T) {

	ctx := context.Background()

	body, err := load_feature(ctx)

	if err != nil {
		t.Fatalf("Failed to load feature, %v", err)
	}

	wr, err := go_writer.NewWriter(ctx, "null://")

	if err != nil {
		t.Fatalf("Failed to create new writer, %v", err)
	}

	id, err := WriteBytes(ctx, wr, body)

	if err != nil {
		t.Fatalf("Failed to write bytes, %v", err)
	}

	if id != 101736545 {
		t.Fatalf("Unexpected ID returned: %d", id)
	}

}

func load_feature(ctx context.Context) ([]byte, error) {

	cwd, err := os.Getwd()

	if err != nil {
		return nil, fmt.Errorf("Failed to determine current working directory, %w", err)
	}

	fixtures := filepath.Join(cwd, "fixtures")
	feature_path := filepath.Join(fixtures, "101736545.geojson")

	fh, err := os.Open(feature_path)

	if err != nil {
		return nil, fmt.Errorf("Failed to open %s, %w", feature_path, err)
	}

	defer fh.Close()

	body, err := io.ReadAll(fh)

	if err != nil {
		return nil, fmt.Errorf("Failed to read %s, %w", feature_path, err)
	}

	return body, nil
}
