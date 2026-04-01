package iterator_test

import (
	"fmt"
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-toolbox/iterator"
	"github.com/KrischanCS/go-toolbox/tuple"
)

func ExampleUnique() {
	i := iterator.Of(1, 1, 2, 1, 2, 3, 3, 1)

	for v := range iterator.Unique(i) {
		fmt.Println(v)
	}

	// Output:
	// 1
	// 2
	// 3
}

func TestUnique(t *testing.T) {
	t.Parallel()

	type test struct {
		name  string
		input iter.Seq[int]
		want  []int
	}

	tests := []test{
		{
			name:  "Should yield all values if all are different",
			input: iterator.Of(1, 2, 3, 4, 5),
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "Should yield only first value if all are the same",
			input: iterator.Of(1, 1, 1, 1, 1),
			want:  []int{1},
		},
		{
			name:  "Should yield first value of each different value",
			input: iterator.Of(1, 1, 1, 2, 2, 3, 3, 3, 3, 5, 5, 5, 5, 5, 5),
			want:  []int{1, 2, 3, 5},
		},
		{
			name:  "Should yield all in the order of their first appearance",
			input: iterator.Of(1, 2, 1, 2, 2, 3, 1, 2, 3, 2, 1, 4, 4, 4, 1, 2, 1, 3, 4, 5, 3, 2, 1, 2, 3, 4, 1, 1, 1),
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "Should yield nothing if input is empty",
			input: iterator.Of[int](),
			want:  []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			got := make([]int, 0, 16)

			// Act
			for v := range iterator.Unique(tc.input) {
				got = append(got, v)
			}

			// Assert
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestUnique_MustStopOnBreak(t *testing.T) {
	t.Parallel()

	// Arrange
	input := iterator.Of(1, 1, 2, 2, 3, 4, 5)
	breakAt := 3
	got := make([]int, 0, 3)

	// Act

	for v := range iterator.Unique(input) {
		got = append(got, v)

		if v == breakAt {
			break
		}
	}

	// Assert
	want := []int{1, 2, 3}
	assert.Equal(t, want, got)
}

func ExampleUniqueBy() {
	pairs := iterator.Of(
		tuple.PairOf(1, "one"),
		tuple.PairOf(1, "eins"),
		tuple.PairOf(2, "two"),
		tuple.PairOf(2, "zwei"),
	)

	first := func(t tuple.Pair[int, string]) int {
		return t.First()
	}

	for p := range iterator.UniqueBy(pairs, first) {
		fmt.Printf("%d: %s\n", p.First(), p.Second())
	}

	// Output:
	// 1: one
	// 2: two
}

func TestUniqueBy(t *testing.T) {
	t.Parallel()

	type test struct {
		name          string
		input         iter.Seq[tuple.Pair[int, string]]
		getComparable func(tuple.Pair[int, string]) int
		want          []tuple.Pair[int, string]
	}

	tests := []test{
		{
			name: "Should yield all values if all are different",
			input: iterator.Of(
				tuple.PairOf(1, "one"),
				tuple.PairOf(2, "two"),
				tuple.PairOf(3, "three"),
			),
			getComparable: func(t tuple.Pair[int, string]) int {
				return t.First()
			},
			want: []tuple.Pair[int, string]{
				tuple.PairOf(1, "one"),
				tuple.PairOf(2, "two"),
				tuple.PairOf(3, "three"),
			},
		},
		{
			name: "Should yield only first value if all are the same",
			input: iterator.Of(
				tuple.PairOf(1, "one"),
				tuple.PairOf(1, "eins"),
				tuple.PairOf(1, "uno"),
			),
			getComparable: func(t tuple.Pair[int, string]) int {
				return t.First()
			},
			want: []tuple.Pair[int, string]{
				tuple.PairOf(1, "one"),
			},
		},
		{
			name: "Should yield first value of each different value",
			input: iterator.Of(
				tuple.PairOf(1, "one"),
				tuple.PairOf(1, "eins"),
				tuple.PairOf(2, "two"),
				tuple.PairOf(2, "zwei"),
				tuple.PairOf(3, "three"),
			),
			getComparable: func(t tuple.Pair[int, string]) int {
				return t.First()
			},
			want: []tuple.Pair[int, string]{
				tuple.PairOf(1, "one"),
				tuple.PairOf(2, "two"),
				tuple.PairOf(3, "three"),
			},
		},
		{
			name:  "Should yield nothing if input is empty",
			input: iterator.Of[tuple.Pair[int, string]](),
			getComparable: func(t tuple.Pair[int, string]) int {
				return t.First()
			},
			want: []tuple.Pair[int, string]{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			got := make([]tuple.Pair[int, string], 0, 16)

			// Act
			for v := range iterator.UniqueBy(tc.input, tc.getComparable) {
				got = append(got, v)
			}

			// Assert
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestUniqueBy_MustStopOnBreak(t *testing.T) {
	t.Parallel()

	// Arrange
	input := iterator.Of(
		tuple.PairOf(1, "one"),
		tuple.PairOf(1, "eins"),
		tuple.PairOf(2, "two"),
		tuple.PairOf(2, "zwei"),
		tuple.PairOf(3, "three"),
		tuple.PairOf(3, "drei"),
	)
	first := func(t tuple.Pair[int, string]) int {
		return t.First()
	}
	breakAt := 3
	got := make([]tuple.Pair[int, string], 0, 2)

	// Act
	for v := range iterator.UniqueBy(input, first) {
		if v.First() == breakAt {
			break
		}
		got = append(got, v)
	}

	// Assert
	want := []tuple.Pair[int, string]{
		tuple.PairOf(1, "one"),
		tuple.PairOf(2, "two"),
	}
	assert.Equal(t, want, got)
}
