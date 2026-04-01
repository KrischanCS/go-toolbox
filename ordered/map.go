// Package ordered provides ordered map and set implementations that maintain
// insertion order.
//
// They provide O(1) time complexity operations while allowing iterating over
// the elements in the order they were added.
package ordered

import (
	"fmt"
	"iter"
	"strings"
)

// Map is an interface for an ordered map that maintains insertion order.
type Map[K comparable, V any] interface {
	// Add inserts a key-value pair into the map. If the key already exists,
	// the value is updated but the order is preserved.
	Add(key K, value V)

	// Get returns the value associated with the key and a boolean indicating
	// whether the key exists in the map.
	Get(key K) (value V, ok bool)

	// Delete removes a key-value pair from the map and returns the deleted
	// value and a boolean indicating whether the key existed.
	Delete(key K) (value V, ok bool)

	// Contains returns true if the key exists in the map.
	Contains(key K) bool

	// Len returns the number of key-value pairs in the map.
	Len() int

	// All returns an iterator over all key-value pairs in insertion order.
	All() iter.Seq2[K, V]

	fmt.Stringer
}

// NewMap creates a new ordered map with the given initial capacity.
func NewMap[K comparable, V comparable](cap int) Map[K, V] {
	return &oMap[K, V]{
		m: make(map[K]*linkedEntry[K, V], cap),
	}
}

type oMap[K comparable, V comparable] struct {
	m           map[K]*linkedEntry[K, V]
	first, last *linkedEntry[K, V]
}

type linkedEntry[K comparable, V comparable] struct {
	key        K
	value      V
	prev, next *linkedEntry[K, V]
}

// Add inserts a key-value pair into the map. If the key already exists,
// the value is updated but the order is preserved.
func (o *oMap[K, V]) Add(key K, value V) {
	if existing, ok := o.m[key]; ok {
		existing.value = value
		return
	}

	entry := &linkedEntry[K, V]{
		key:   key,
		value: value,
		prev:  o.last,
	}

	o.m[key] = entry

	if o.first == nil {
		o.first = entry
	}

	if o.last != nil {
		o.last.next = entry
	}

	o.last = entry
}

// Get returns the value associated with the key and a boolean indicating
// whether the key exists in the map.
func (o *oMap[K, V]) Get(key K) (v V, ok bool) {
	entry, ok := o.m[key]
	if !ok {
		return
	}

	return entry.value, true
}

// Delete removes a key-value pair from the map and returns the deleted
// value and a boolean indicating whether the key existed.
func (o *oMap[K, V]) Delete(key K) (v V, ok bool) {
	entry, ok := o.m[key]
	if !ok {
		return
	}

	if entry == o.first {
		o.first = entry.next
	} else {
		entry.prev.next = entry.next
	}

	if entry == o.last {
		o.last = entry.prev
	} else {
		entry.next.prev = entry.prev
	}

	delete(o.m, key)

	return entry.value, true
}

// Contains returns true if the key exists in the map.
func (o *oMap[K, V]) Contains(key K) bool {
	_, ok := o.m[key]
	return ok
}

// Len returns the number of key-value pairs in the map.
func (o *oMap[K, V]) Len() int {
	return len(o.m)
}

// All returns an iterator over all key-value pairs in insertion order.
func (o *oMap[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		current := o.first
		for current != nil {
			if !yield(current.key, current.value) {
				return
			}
			current = current.next
		}
	}
}

// String returns a string representation of the ordered map in the format:
// "OrderedMap[{key1: value1}, {key2: value2}, ...]"
func (o *oMap[K, V]) String() string {
	var sb strings.Builder
	sb.WriteString("(OrderedMap[")

	first := true
	for k, v := range o.All() {
		if !first {
			sb.WriteString(", ")
		}
		first = false
		sb.WriteString(fmt.Sprintf("{%v: %v}", k, v))
	}

	sb.WriteString("])")
	return sb.String()
}
