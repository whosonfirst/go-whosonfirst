// Package feature defines errors for accessing Who's On First feature properties.
package feature

import (
	"fmt"
)

// PropertyNotFoundErr identifies a property that is absent from a feature.
type PropertyNotFoundErr struct {
	property string
}

// Error returns a message naming the missing property.
func (e *PropertyNotFoundErr) Error() string {
	return fmt.Sprintf("'%s' property not found", e.property)
}

// Property returns the name of the missing property.
func (e *PropertyNotFoundErr) Property() string {
	return e.property
}

// PropertyNotFoundError constructs an error for the specified missing property.
func PropertyNotFoundError(prop string) *PropertyNotFoundErr {
	return &PropertyNotFoundErr{
		property: prop,
	}
}

// IsPropertyNotFoundError reports whether target is a *PropertyNotFoundErr. It does not unwrap errors.
func IsPropertyNotFoundError(target error) bool {
	switch target.(type) {
	case *PropertyNotFoundErr:
		return true
	default:
		return false
	}
}
