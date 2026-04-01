// Package optional provides a generic Optional type for Go.
//
// For a lot of use cases, I don't like mixing the concepts of referencing and
// optionality, as it happens with pointers, so this provides an alternative.
//
// The Optionals implemented in this package handle marshalling and unmarshalling
// to/from json and xml and are handling null-values, not existing values and
// omitzero/omitempty.
package optional

import "fmt"

// Optional represents a value of type T that may or may not be present.
type Optional[T any] struct {
	value   T
	present bool
}

// Of creates a new present Optional with the given value.
func Of[T any](value T) Optional[T] {
	return Optional[T]{value: value, present: true}
}

// Empty creates a new empty Optional.
func Empty[T any]() Optional[T] {
	return Optional[T]{present: false}
}

// Get returns the value if present, and a boolean indicating its presence.
func (o Optional[T]) Get() (value T, ok bool) {
	return o.value, o.present
}

// Or returns the value if present, or the provided fallback value if empty.
func (o Optional[T]) Or(fallback T) T {
	if !o.present {
		return fallback
	}

	return o.value
}

// IfOk executes the provided function with the value if it is present.
func (o Optional[T]) IfOk(fn func(T)) {
	if o.present {
		fn(o.value)
	}
}

// String returns a string representation in the format:
//   - If Present: Optional[{{type}}]: {{value}}
//   - If Empty: Optional[{{type}}]: <empty>
func (o Optional[T]) String() string {
	if !o.present {
		return fmt.Sprintf("(Optional[%T] <empty>)", o.value)
	}

	return fmt.Sprintf("(Optional[%T]: %v)", o.value, o.value)
}
