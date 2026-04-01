package stack_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-toolbox/stack"
)

func TestOf(t *testing.T) {
	type test struct {
		name   string
		values []int
		want   []int
	}

	tests := []test{
		{
			name:   "Should create an empty stack when no values are given",
			values: []int{},
			want:   []int{},
		},
		{
			name:   "Should create a stack with one value when one value is given",
			values: []int{42},
			want:   []int{42},
		},
		{
			name:   "Should create a stack with multiple values where the last one is on top when multiple values are given",
			values: []int{1, 2, 3, 4, 5},
			want:   []int{5, 4, 3, 2, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			intStack := stack.Of(tc.values...)

			// Assert

			for _, v := range tc.want {
				got, _ := intStack.Pop()
				assert.Equal(t, v, got)
			}
		})
	}
}

func TestWithCapacity(t *testing.T) {
	type test struct {
		name         string
		capacity     int
		wantCapacity int
	}

	tests := []test{
		{
			name:         "Should create an empty stack with zero capacity",
			capacity:     0,
			wantCapacity: 0,
		},
		{
			name:         "Should create an empty stack with a positive capacity",
			capacity:     10,
			wantCapacity: 10,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			intStack := stack.WithCapacity[int](tc.capacity)

			// Assert
			assert.Equal(t, tc.wantCapacity, intStack.Cap())
			assert.Equal(t, 0, intStack.Len())
		})
	}
}

func TestStack_Push(t *testing.T) {
	type test struct {
		name       string
		initValues []int
		push       []int
		want       []int
	}

	tests := []test{
		{
			name:       "Should not modify the stack when no values are pushed",
			initValues: []int{1, 2, 3},
			push:       []int{},
			want:       []int{3, 2, 1},
		},
		{
			name:       "Should push a value onto an empty stack",
			initValues: []int{},
			push:       []int{42},
			want:       []int{42},
		},
		{
			name:       "Should push a single value onto an empty stack",
			initValues: []int{1, 2, 3},
			push:       []int{4},
			want:       []int{4, 3, 2, 1},
		},
		{
			name:       "Should push multiple values onto an empty stack",
			initValues: []int{1, 2, 3},
			push:       []int{4, 5},
			want:       []int{5, 4, 3, 2, 1},
		},
		{
			name:       "Should push multiple values onto a non-empty stack",
			initValues: []int{1, 2, 3},
			push:       []int{4, 5, 6},
			want:       []int{6, 5, 4, 3, 2, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			intStack := stack.Of(tc.initValues...)

			// Act
			for _, value := range tc.push {
				intStack.Push(value)
			}

			// Assert
			for _, v := range tc.want {
				got, _ := intStack.Pop()
				assert.Equal(t, v, got)
			}

			assert.True(t, intStack.IsEmpty())
		})
	}
}

func TestStack_Pop(t *testing.T) {
	type test struct {
		name   string
		values []int
		want   []int
		ok     bool
	}

	tests := []test{
		{
			name:   "Should return false and zero value when popping from an empty stack",
			values: []int{},
			want:   []int{},
			ok:     false,
		},
		{
			name:   "Should pop the last value from a stack with one value",
			values: []int{42},
			want:   []int{42},
			ok:     true,
		},
		{
			name:   "Should pop all values from a stack with multiple values",
			values: []int{1, 2, 3, 4, 5},
			want:   []int{5, 4, 3, 2, 1},
			ok:     true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			intStack := stack.Of(tc.values...)

			// Act & Assert
			for _, v := range tc.want {
				got, ok := intStack.Pop()
				assert.Equal(t, v, got)
				assert.Equal(t, true, ok)
			}

			got, ok := intStack.Pop()
			assert.Equal(t, 0, got)
			assert.Equal(t, false, ok)
		})
	}
}

func TestStack_PopN(t *testing.T) {
	type test struct {
		name   string
		values []int
		n      int
		want   []int
	}

	tests := []test{
		{
			name:   "Should return an empty slice when popping from an empty stack",
			values: []int{},
			n:      3,
			want:   []int{},
		},
		{
			name:   "Should pop one value from a stack with one value",
			values: []int{42},
			n:      1,
			want:   []int{42},
		},
		{
			name:   "Should pop all values from a stack with multiple values",
			values: []int{1, 2, 3, 4, 5},
			n:      5,
			want:   []int{5, 4, 3, 2, 1},
		},
		{
			name:   "Should pop the top n values from a stack with multiple values",
			values: []int{1, 2, 3, 4, 5},
			n:      3,
			want:   []int{5, 4, 3},
		},
		{
			name:   "Should pop all values when n is greater than the stack size",
			values: []int{1, 2, 3, 4, 5},
			n:      10,
			want:   []int{5, 4, 3, 2, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			intStack := stack.Of(tc.values...)

			// Act
			got := intStack.PopN(tc.n)

			// Assert
			assert.Equal(t, tc.want, got)
			if tc.n >= len(tc.values) {
				assert.True(t, intStack.IsEmpty())
			} else {
				assert.Equal(t, len(tc.values)-tc.n, intStack.Len())
			}
		})
	}
}

func TestStack_PopAll(t *testing.T) {
	type test struct {
		name   string
		values []int
		want   []int
	}

	tests := []test{
		{
			name:   "Should return an empty slice when popping all from an empty stack",
			values: []int{},
			want:   []int{},
		},
		{
			name:   "Should pop all values from a stack with one value",
			values: []int{42},
			want:   []int{42},
		},
		{
			name:   "Should pop all values from a stack with multiple values",
			values: []int{1, 2, 3, 4, 5},
			want:   []int{5, 4, 3, 2, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			intStack := stack.Of(tc.values...)

			// Act
			got := intStack.PopAll()

			// Assert
			assert.Equal(t, tc.want, got)
			assert.True(t, intStack.IsEmpty())
		})
	}
}

func TestStack_Peek(t *testing.T) {
	type test struct {
		name   string
		values []int
		want   int
		ok     bool
	}

	tests := []test{
		{
			name:   "Should return zero value and false when peeking an empty stack",
			values: []int{},
			want:   0,
			ok:     false,
		},
		{
			name:   "Should peek the last value from a stack with one value",
			values: []int{42},
			want:   42,
			ok:     true,
		},
		{
			name:   "Should peek the top value from a stack with multiple values",
			values: []int{1, 2, 3, 4, 5},
			want:   5,
			ok:     true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			intStack := stack.Of(tc.values...)

			// Act
			got, ok := intStack.Peek()

			// Assert
			assert.Equal(t, tc.want, got)
			assert.Equal(t, tc.ok, ok)
			assert.Equal(t, len(tc.values), intStack.Len(), "Peek should not modify the stack")
		})
	}
}

func TestStack_PeekN(t *testing.T) {
	type test struct {
		name   string
		values []int
		n      int
		want   []int
	}

	tests := []test{
		{
			name:   "Should return an empty slice when peeking an empty stack",
			values: []int{},
			n:      3,
			want:   []int{},
		},
		{
			name:   "Should return one value from a stack with one value",
			values: []int{42},
			n:      1,
			want:   []int{42},
		},
		{
			name:   "Should return 3 values from a stack with 5 values, when PeekN is called with n=3",
			values: []int{1, 2, 3, 4, 5},
			n:      3,
			want:   []int{5, 4, 3},
		},
		{
			name:   "Should return all values if n is greater than stack size",
			values: []int{1, 2, 3, 4, 5},
			n:      8,
			want:   []int{5, 4, 3, 2, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			intStack := stack.Of(tc.values...)

			// Act
			got := intStack.PeekN(tc.n)

			// Assert
			assert.Equal(t, tc.want, got)
			assert.Equal(t, len(tc.values), intStack.Len(), "PeekN should not modify the stack")
		})
	}
}

func TestStack_PeekAll(t *testing.T) {
	type test struct {
		name   string
		values []int
		want   []int
	}

	tests := []test{
		{
			name:   "Should return an empty slice when peeking all from an empty stack",
			values: []int{},
			want:   []int{},
		},
		{
			name:   "Should peek all values from a stack with one value",
			values: []int{42},
			want:   []int{42},
		},
		{
			name:   "Should peek all values from a stack with multiple values",
			values: []int{1, 2, 3, 4, 5},
			want:   []int{5, 4, 3, 2, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			intStack := stack.Of(tc.values...)

			// Act
			got := intStack.PeekAll()

			// Assert
			assert.Equal(t, tc.want, got)
			assert.Equal(t, len(tc.values), intStack.Len(), "PeekAll should not modify the stack")
		})
	}
}

func TestStack_Grow(t *testing.T) {
	s := stack.Of(1, 2, 3)

	assert.Equal(t, 3, s.Cap())

	s.Grow(3)

	assert.GreaterOrEqual(t, s.Cap(), 6)

	s.Grow(5)

	assert.GreaterOrEqual(t, s.Cap(), 9)
}

func TestStack_Clear(t *testing.T) {
	type test struct {
		name   string
		values []int
	}

	tests := []test{
		{
			name:   "Should clear an empty stack",
			values: []int{},
		},
		{
			name:   "Should clear a stack with one value",
			values: []int{42},
		},
		{
			name:   "Should clear a stack with multiple values",
			values: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			intStack := stack.Of(tc.values...)
			capacityBefore := intStack.Cap()

			// Act
			intStack.Clear()

			// Assert
			assert.True(t, intStack.IsEmpty())
			assert.Equal(t, 0, intStack.Len())
			assert.Equal(t, capacityBefore, intStack.Cap(), "Clearing the stack should not change its capacity")
		})
	}
}

func TestStack_String(t *testing.T) {
	type test struct {
		name   string
		values []int
		want   string
	}

	tests := []test{
		{
			name:   "Empty",
			values: []int{},
			want:   "(Stack[int]: <empty>)",
		},
		{
			name:   "Single value",
			values: []int{42},
			want:   "(Stack[int]: Top< 42 |Bottom)",
		},
		{
			name:   "Multiple values",
			values: []int{1, 2, 3, 4, 5},
			want:   "(Stack[int]: Top< 5, 4, 3, 2, 1 |Bottom)",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			intStack := stack.Of(tc.values...)

			// Act
			got := intStack.String()

			// Assert
			assert.Equal(t, tc.want, got)
		})
	}
}
