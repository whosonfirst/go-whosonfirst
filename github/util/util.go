// Package util provides a legacy helper for constructing a GitHub API client.
package util

import (
	"context"

	"github.com/google/go-github/v88/github"
	"github.com/whosonfirst/go-whosonfirst/v4/github/client"
)

// NewClientAndContext constructs a GitHub client and returns it with a background context.
//
// Deprecated: Use [client.NewClient] with a caller-provided context.
func NewClientAndContext(token string) (*github.Client, context.Context, error) {

	ctx := context.Background()

	cl, err := client.NewClient(ctx, token)

	if err != nil {
		return nil, nil, err
	}

	return cl, ctx, nil
}
