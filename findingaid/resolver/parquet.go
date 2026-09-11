package resolver

import (
	"context"
	"fmt"
	"net/url"

	parquet "github.com/whosonfirst/go-whosonfirst/v4/findingaid/parquet"
)

// type ParquetResolver implements the `Resolver` interface for data stored in a Parquet database
// produced by `github.com/whosonfirst/go-whosonfirst/v4/findingaid/parquet.ParquetProducer`
type ParquetResolver struct {
	Resolver
	parquet_uri string
}

func init() {
	ctx := context.Background()
	RegisterResolver(ctx, "parquet", NewParquetResolver)
}

// NewParquetResolver will return a new `Resolver` instance for resolving repository names
// and IDs stored in a Parquet database.
func NewParquetResolver(ctx context.Context, uri string) (Resolver, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse URL, %w", err)
	}

	parquet_uri := u.Path

	q := u.Query()

	if q.Has("parquet-uri") {
		parquet_uri = q.Get("parquet-uri")
	}

	f := &ParquetResolver{
		parquet_uri: parquet_uri,
	}

	return f, nil
}

// GetRepo returns the name of the repository associated with this ID in a Who's On First finding aid.
func (r *ParquetResolver) GetRepo(ctx context.Context, id int64) (string, error) {

	return parquet.GetRepo(r.parquet_uri, id)
}

func (r ParquetResolver) Close() error {
	return nil
}
