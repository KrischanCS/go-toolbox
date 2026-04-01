// Package stack provides a simple, generic stack implementation in go.
package stack

import (
	"fmt"
	"slices"
	"strings"
)

// Stack is a generic stack type that can hold elements of any type T.
type Stack[T any] struct {
	s []T
}

// Of creates a new Stack with the provided values. The last value in the argument list is the top
// of the stack.
func Of[T any](values ...T) Stack[T] {
	return Stack[T]{
		s: slices.Clone(values),
	}
}

// WithCapacity creates an empty Stack with the specified initial capacity.
func WithCapacity[T any](capacity int) Stack[T] {
	return Stack[T]{make([]T, 0, capacity)}
}

// Push adds a new element to the top of the stack.
func (s *Stack[T]) Push(value T) {
	s.s = append(s.s, value)
}

// Pop removes and returns the top element of the stack.
//
// If the stack is empty, ok will be false and the zero value of T will be returned.
func (s *Stack[T]) Pop() (value T, ok bool) {
	if len(s.s) == 0 {
		return value, false
	}

	iLast := len(s.s) - 1

	v := s.s[iLast]
	s.s = s.s[:iLast]

	return v, true
}

// PopN removes and returns the top n elements of the stack as a slice, the top element first.
//
// If n is greater than the number of elements in the stack, it returns all elements.
func (s *Stack[T]) PopN(n int) []T {
	l := min(n, s.Len())
	values := make([]T, l)

	for i := range l {
		values[i], _ = s.Pop()
	}

	return values
}

// PopAll removes and returns all elements of the stack as a slice, the top element first.
func (s *Stack[T]) PopAll() []T {
	values := make([]T, s.Len())
	for i := range values {
		values[i], _ = s.Pop()
	}

	return values
}

// Peek returns the top element of the stack without removing it.
func (s *Stack[T]) Peek() (value T, ok bool) {
	if len(s.s) == 0 {
		return value, false
	}

	return s.s[len(s.s)-1], true
}

// PeekN returns the top n elements of the stack without removing them, the top element first.
func (s *Stack[T]) PeekN(n int) []T {
	sLen := s.Len()

	l := min(n, sLen)
	values := make([]T, l)

	for i := range l {
		values[i] = s.s[sLen-1-i]
	}

	return values
}

// PeekAll returns all elements of the stack without removing them, the top element first.
func (s *Stack[T]) PeekAll() []T {
	values := make([]T, s.Len())
	for i := len(s.s) - 1; i >= 0; i-- {
		values[i] = s.s[len(s.s)-i-1]
	}

	return values
}

// Len returns the number of elements in the stack.
func (s *Stack[T]) Len() int {
	return len(s.s)
}

// Cap returns the current capacity of the stack.
func (s *Stack[T]) Cap() int {
	return cap(s.s)
}

// Grow increases the stacks capacity, if necessary, to guarantee space for another n elements.
//
// After Grow(n), at least n elements can be pushed to the stack without another allocation.
//
// If n is negative or too large to allocate the memory, Grow panics.
func (s *Stack[T]) Grow(capacity int) {
	s.s = slices.Grow(s.s, capacity)
}

// IsEmpty checks if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.s) == 0
}

// Clear removes all elements from the stack.
func (s *Stack[T]) Clear() {
	s.s = s.s[:0]
}

// String returns a string representation of the stack in the format:
//
//   - non empty: (Stack[{{type}}]: Top< {{top element}},  ..., {{bottom element}} |Bottom)
//   - empty: (Stack[{{type}}]: <empty>)
func (s *Stack[T]) String() string {
	var tmp T

	if len(s.s) == 0 {
		return fmt.Sprintf("(Stack[%T]: <empty>)", tmp)
	}

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("(Stack[%T]: Top< ", tmp))
	sb.WriteString(fmt.Sprintf("%v", s.s[len(s.s)-1]))
	for _, v := range slices.Backward(s.s[:len(s.s)-1]) {
		sb.WriteString(fmt.Sprintf(", %v", v))
	}
	sb.WriteString(fmt.Sprintf(" |Bottom)"))

	return sb.String()
}
