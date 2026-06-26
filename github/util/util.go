package util

import (
	"context"

	"github.com/google/go-github/v88/github"
	"github.com/whosonfirst/go-whosonfirst/v4/github/client"
)

// This method is deprecated
func NewClientAndContext(token string) (*github.Client, context.Context, error) {

	ctx := context.Background()

	cl, err := client.NewClient(ctx, token)

	if err != nil {
		return nil, nil, err
	}

	return cl, ctx, nil
}
