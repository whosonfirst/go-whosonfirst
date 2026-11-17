// Package derivatives defines sentinel errors for unimplemented operations and missing results.
package derivatives

import (
	"errors"
)

// ErrNotImplemented indicates that the requested operation is not implemented.
var ErrNotImplemented = errors.New("Not implemented")

// ErrNotFound indicates that the requested result was not found.
var ErrNotFound = errors.New("Not found")
