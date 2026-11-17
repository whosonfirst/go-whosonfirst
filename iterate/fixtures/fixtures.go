// Package fixtures provides fixtures for testing record iterators.
package fixtures

import (
	"embed"
)

//go:embed data/*
var FS embed.FS
