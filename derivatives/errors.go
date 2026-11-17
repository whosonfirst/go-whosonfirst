// Package derivatives defines errors shared by the Who's On First derivative generators.
package derivatives

import (
	"errors"
)

var ErrNotImplemented = errors.New("Not implemented")

var ErrNotFound = errors.New("Not found")
