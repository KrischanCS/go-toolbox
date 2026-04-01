package ordered

import (
	"iter"
)

type Map[K comparable, V any] interface {
	Add(key K, value V)
	Delete(key K) (value V, ok bool)

	Contains(key K) bool
	Len() int

	All() iter.Seq2[K, V]
}

func NewOrdered[K comparable, V any](cap int) Map[K, V] {
	return &ordered[K, V]{
		m: make(map[K]*linkedEntry[K, V], cap),
	}
}

type ordered[K comparable, V any] struct {
	m           map[K]*linkedEntry[K, V]
	first, last *linkedEntry[K, V]
}

func (o *ordered[K, V]) Add(key K, value V) {
	entry := &linkedEntry[K, V]{
		key:   key,
		value: value,
		prev:  o.last,
	}

	if o.first == nil {
		o.first = entry
	}

	if o.last != nil {
		o.last.next = entry
	}

	o.last = entry
}

func (o *ordered[K, V]) Delete(key K) (v V, ok bool) {
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
	}

	delete(o.m, key)

	return entry.value, true
}

func (o *ordered[K, V]) Contains(key K) bool {
	_, ok := o.m[key]
	return ok
}

func (o *ordered[K, V]) Len() int {
	return len(o.m)
}

func (o *ordered[K, V]) All() iter.Seq2[K, V] {
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

func ContainsValue[K comparable, V comparable](m Map[K, V], value V) bool {
	for _, v := range m.All() {
		if value == v {
			return true
		}
	}

	return false
}

type linkedEntry[K comparable, V any] struct {
	key        K
	value      V
	prev, next *linkedEntry[K, V]
}
