package document

import (
	"context"
)

// PrepareDocumentFunc is a common method signature updating a Who's On First document for indexing in OpenSearch.
type PrepareDocumentFunc func(context.Context, []byte) ([]byte, error)
