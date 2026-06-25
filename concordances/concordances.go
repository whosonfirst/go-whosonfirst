package concordances

import (
	"context"
	"fmt"
	"io"
	"sort"
	"sync"

	"github.com/whosonfirst/go-whosonfirst/v4/feature/properties"
	"github.com/whosonfirst/go-whosonfirst/v4/iterate"
)

// ListKeys() returns the list of unique keys for all the concordances found in 'iterator_sources'.
// 'iterator_uri' is expected to be a valid `whosonfirst/go-whosonfirst-iterate/v2` URI and 'iterator_sources'
// a list of URIs to be crawled.
func ListKeys(ctx context.Context, iterator_uri string, iterator_sources ...string) ([]string, error) {

	sources := new(sync.Map)

	iter, err := iterate.NewIterator(ctx, iterator_uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to create iterator, %w", err)
	}

	for rec, err := range iter.Iterate(ctx, iterator_sources...) {

		if err != nil {
			return nil, err
		}

		defer rec.Body.Close()

		body, err := io.ReadAll(rec.Body)

		if err != nil {
			return nil, fmt.Errorf("Failed to read %s, %w", rec.Path, err)
		}

		c := properties.Concordances(body)

		for src, _ := range c {
			sources.Store(src, true)
		}

	}

	concordances_keys := make([]string, 0)

	sources.Range(func(k interface{}, v interface{}) bool {
		concordances_keys = append(concordances_keys, k.(string))
		return true
	})

	sort.Strings(concordances_keys)
	return concordances_keys, nil
}
