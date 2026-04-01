package iterator_test

import (
	"fmt"
	"iter"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-toolbox/iterator"
	"github.com/KrischanCS/go-toolbox/tuple"
)

func ExampleZip() {
	numbers := iterator.Of(1, 2, 3, 4)
	letters := iterator.Of("a", "b", "c")

	for pair := range iterator.Zip[int, string](numbers, letters) {
		fmt.Println(pair.First(), pair.Second())
	}

	// Output:
	// 1 a
	// 2 b
	// 3 c
}

//nolint:funlen
func TestZip(t *testing.T) {
	t.Parallel()

	type testCase[L, R any] struct {
		name       string
		leftInput  iter.Seq[L]
		rightInput iter.Seq[R]
		want       []tuple.Pair[L, R]
	}

	testCases := []testCase[int, string]{
		{
			name:       "empty slices",
			leftInput:  iterator.Of[int](),
			rightInput: iterator.Of[string](),
			want:       []tuple.Pair[int, string]{},
		},
		{
			name:       "one element",
			leftInput:  iterator.Of(1),
			rightInput: iterator.Of("a"),
			want: []tuple.Pair[int, string]{
				tuple.PairOf(1, "a"),
			},
		},
		{
			name:       "multiple elements",
			leftInput:  iterator.Of(1, 2, 3),
			rightInput: iterator.Of("a", "b", "c"),
			want: []tuple.Pair[int, string]{
				tuple.PairOf(1, "a"),
				tuple.PairOf(2, "b"),
				tuple.PairOf(3, "c"),
			},
		},
		{
			name:       "len(leftInput) < len(rightInput)",
			leftInput:  iterator.Of(1, 2, 3),
			rightInput: iterator.Of("a", "b"),
			want: []tuple.Pair[int, string]{
				tuple.PairOf(1, "a"),
				tuple.PairOf(2, "b"),
			},
		},
		{
			name:       "len(leftInput) > len(rightInput)",
			leftInput:  iterator.Of(1, 2),
			rightInput: iterator.Of("a", "b", "c"),
			want: []tuple.Pair[int, string]{
				tuple.PairOf(1, "a"),
				tuple.PairOf(2, "b"),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			got := make([]tuple.Pair[int, string], 0, len(tc.want))

			// Act
			for p := range iterator.Zip(tc.leftInput, tc.rightInput) {
				got = append(got, p)
			}

			// Assert
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestZip_withBreak(t *testing.T) {
	t.Parallel()

	// Arrange
	l := iterator.Of(1, 2, 3, 4, 5, 6)

	r := iterator.Of("a", "b", "c", "d", "e", "f")

	got := make([]tuple.Pair[int, string], 0, 2)

	stop := tuple.PairOf[int, string](3, "c")

	// Act
	for e := range iterator.Zip(l, r) {
		if e == stop {
			break
		}

		got = append(got, e)
	}

	// Assert
	want := []tuple.Pair[int, string]{
		tuple.PairOf(1, "a"),
		tuple.PairOf(2, "b"),
	}
	assert.Equal(t, want, got)
}

func ExampleZip3() {
	numbers := iterator.Of(1, 2, 3, 4)
	letters := iterator.Of("a", "b", "c")
	bytes := iterator.Of('!', '@', '#')

	for triple := range iterator.Zip3[int, string, rune](numbers, letters, bytes) {
		fmt.Printf("%d %s %c\n", triple.First(), triple.Second(), triple.Third())
	}

	// Output:
	// 1 a !
	// 2 b @
	// 3 c #
}

func TestZip3(t *testing.T) {
	t.Parallel()

	type testCase[A, B, C any] struct {
		name   string
		inputA iter.Seq[A]
		inputB iter.Seq[B]
		inputC iter.Seq[C]
		want   []tuple.Triple[A, B, C]
	}

	testCases := []testCase[int, string, rune]{
		{
			name:   "empty slices",
			inputA: iterator.Of[int](),
			inputB: iterator.Of[string](),
			inputC: iterator.Of[rune](),
			want:   []tuple.Triple[int, string, rune]{},
		},
		{
			name:   "one element",
			inputA: iterator.Of(1),
			inputB: iterator.Of("a"),
			inputC: iterator.Of('!'),
			want: []tuple.Triple[int, string, rune]{
				tuple.TripleOf(1, "a", '!'),
			},
		},
		{
			name:   "multiple elements",
			inputA: iterator.Of(1, 2, 3),
			inputB: iterator.Of("a", "b", "c"),
			inputC: iterator.Of('!', '@', '#'),
			want: []tuple.Triple[int, string, rune]{
				tuple.TripleOf(1, "a", '!'),
				tuple.TripleOf(2, "b", '@'),
				tuple.TripleOf(3, "c", '#'),
			},
		},
		{
			name:   "len(inputA) < len(inputB) and len(inputA) < len(inputC)",
			inputA: iterator.Of(1, 2, 3),
			inputB: iterator.Of("a", "b"),
			inputC: iterator.Of('!', '@', '#', '$'),
			want: []tuple.Triple[int, string, rune]{
				tuple.TripleOf(1, "a", '!'),
				tuple.TripleOf(2, "b", '@'),
			},
		},
		{
			name:   "len(inputB) < len(inputA) and len(inputB) < len(inputC)",
			inputA: iterator.Of(1, 2, 3, 4),
			inputB: iterator.Of("a", "b"),
			inputC: iterator.Of('!', '@', '#'),
			want: []tuple.Triple[int, string, rune]{
				tuple.TripleOf(1, "a", '!'),
				tuple.TripleOf(2, "b", '@'),
			},
		},
		{
			name:   "len(inputC) < len(inputA) and len(inputC) < len(inputB)",
			inputA: iterator.Of(1, 2, 3, 4),
			inputB: iterator.Of("a", "b", "c"),
			inputC: iterator.Of('!', '@'),
			want: []tuple.Triple[int, string, rune]{
				tuple.TripleOf(1, "a", '!'),
				tuple.TripleOf(2, "b", '@'),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			got := make([]tuple.Triple[int, string, rune], 0, len(tc.want))

			// Act
			for p := range iterator.Zip3(tc.inputA, tc.inputB, tc.inputC) {
				got = append(got, p)
			}

			// Assert
			assert.Equal(t, tc.want, got)

		})
	}
}

func TestZip3_withBreak(t *testing.T) {
	t.Parallel()

	// Arrange
	inputA := iterator.Of(1, 2, 3, 4, 5, 6)
	inputB := iterator.Of("a", "b", "c", "d", "e", "f")
	inputC := iterator.Of('!', '@', '#')

	got := make([]tuple.Triple[int, string, rune], 0, 2)

	stop := 3

	// Act
	for e := range iterator.Zip3(inputA, inputB, inputC) {
		if e.First() == stop {
			break
		}

		got = append(got, e)
	}

	// Assert
	want := []tuple.Triple[int, string, rune]{
		tuple.TripleOf(1, "a", '!'),
		tuple.TripleOf(2, "b", '@'),
	}
	assert.Equal(t, want, got)
}
