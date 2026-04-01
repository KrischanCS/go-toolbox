package ordered

import (
	"fmt"
	"iter"
	"strings"
)

// Set is an interface for an ordered set that maintains insertion order.
type Set[T comparable] interface {
	// Add inserts a value into the set. If the value already exists,
	// the set is unchanged and the order is preserved.
	Add(value T)

	// Remove removes a value from the set and returns a boolean indicating
	// whether the value existed.
	Remove(value T) bool

	// Contains returns true if the value exists in the set.
	Contains(value T) bool

	// Len returns the number of values in the set.
	Len() int

	// IsEmpty returns true if the set is empty.
	IsEmpty() bool

	// Clear removes all values from the set.
	Clear()

	// Clone creates a shallow copy of the set.
	Clone() Set[T]

	// Values returns a slice of all values in insertion order.
	Values() []T

	// All returns an iterator over all values in insertion order.
	All() iter.Seq[T]

	// String returns a string representation of the ordered set.
	String() string
}

// NewSet creates a new ordered set with the given initial capacity.
func NewSet[T comparable](cap int) Set[T] {
	return &oSet[T]{
		m: make(map[T]*linkedSetEntry[T], cap),
	}
}

// SetOf creates a new ordered set with the given values.
func SetOf[T comparable](values ...T) Set[T] {
	s := &oSet[T]{
		m: make(map[T]*linkedSetEntry[T], len(values)),
	}

	for _, v := range values {
		s.Add(v)
	}

	return s
}

type oSet[T comparable] struct {
	m           map[T]*linkedSetEntry[T]
	first, last *linkedSetEntry[T]
}

type linkedSetEntry[T comparable] struct {
	value      T
	prev, next *linkedSetEntry[T]
}

// Add inserts a value into the set. If the value already exists,
// the set is unchanged and the order is preserved.
func (s *oSet[T]) Add(value T) {
	if _, ok := s.m[value]; ok {
		return
	}

	entry := &linkedSetEntry[T]{
		value: value,
		prev:  s.last,
	}

	s.m[value] = entry

	if s.first == nil {
		s.first = entry
	}

	if s.last != nil {
		s.last.next = entry
	}

	s.last = entry
}

// Remove removes a value from the set and returns a boolean indicating
// whether the value existed.
func (s *oSet[T]) Remove(value T) bool {
	entry, ok := s.m[value]
	if !ok {
		return false
	}

	if entry == s.first {
		s.first = entry.next
	} else {
		entry.prev.next = entry.next
	}

	if entry == s.last {
		s.last = entry.prev
	} else {
		entry.next.prev = entry.prev
	}

	delete(s.m, value)

	return true
}

// Contains returns true if the value exists in the set.
func (s *oSet[T]) Contains(value T) bool {
	_, ok := s.m[value]
	return ok
}

// Len returns the number of values in the set.
func (s *oSet[T]) Len() int {
	return len(s.m)
}

// IsEmpty returns true if the set is empty.
func (s *oSet[T]) IsEmpty() bool {
	return len(s.m) == 0
}

// Clear removes all values from the set.
func (s *oSet[T]) Clear() {
	s.m = make(map[T]*linkedSetEntry[T])
	s.first = nil
	s.last = nil
}

// Clone creates a shallow copy of the set.
func (s *oSet[T]) Clone() Set[T] {
	clone := &oSet[T]{
		m: make(map[T]*linkedSetEntry[T], len(s.m)),
	}

	for v := range s.All() {
		clone.Add(v)
	}

	return clone
}

// Values returns a slice of all values in insertion order.
func (s *oSet[T]) Values() []T {
	values := make([]T, 0, len(s.m))

	for v := range s.All() {
		values = append(values, v)
	}

	return values
}

// All returns an iterator over all values in insertion order.
func (s *oSet[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		current := s.first
		for current != nil {
			if !yield(current.value) {
				return
			}
			current = current.next
		}
	}
}

// String returns a string representation of the ordered set in the format:
// "OrderedSet[value1, value2, ...]"
func (s *oSet[T]) String() string {
	if s.IsEmpty() {
		return fmt.Sprintf("(OrderedSet[%T]: <empty>)", *new(T))
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("(OrderedSet[%T]: [", *new(T)))

	first := true
	for v := range s.All() {
		if !first {
			sb.WriteString(", ")
		}
		first = false
		sb.WriteString(fmt.Sprintf("%v", v))
	}

	sb.WriteString("])")
	return sb.String()
}
