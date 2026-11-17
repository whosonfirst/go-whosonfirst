//go:build !no_centroid

// Package wof assembles the subcommands for the wof command-line application.
package wof

// Experimental: Moving specific imports in to build tags so people to build trimmed-down
// binaries if they want to

import (
	_ "github.com/whosonfirst/go-whosonfirst/v4/wof/centroid"
)
