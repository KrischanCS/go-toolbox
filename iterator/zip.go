package iterator

import (
	"iter"

	"github.com/KrischanCS/go-toolbox/tuple"
)

// Zip creates a new [iter.Seq] from the left and right iterators, which yields
// pairs of values from both of the same iterations.
//
// The resulting iterator will stop when the shorter of the two iterators stops.
func Zip[L, R any](left iter.Seq[L], right iter.Seq[R]) iter.Seq[tuple.Pair[L, R]] {
	return func(yield func(tuple.Pair[L, R]) bool) {
		valuesRight, stop := iter.Pull(right)
		defer stop()

		for valueL := range left {
			valueR, ok := valuesRight()
			if !ok {
				return
			}

			if !yield(tuple.PairOf[L, R](valueL, valueR)) {
				return
			}
		}
	}
}

// Zip3 creates a new [iter.Seq] from the iterators a, b and c, which yields
// triples of values from all three.
//
// The resulting iterator will stop when the shortest of the three iterators stops.
func Zip3[A, B, C any](a iter.Seq[A], b iter.Seq[B], c iter.Seq[C]) iter.Seq[tuple.Triple[A, B, C]] {
	return func(yield func(tuple.Triple[A, B, C]) bool) {
		valuesB, stopB := iter.Pull(b)
		defer stopB()

		valuesC, stopC := iter.Pull(c)
		defer stopC()

		for valueA := range a {
			valueB, okB := valuesB()
			if !okB {
				return
			}

			valueC, okC := valuesC()
			if !okC {
				return
			}

			if !yield(tuple.TripleOf(valueA, valueB, valueC)) {
				return
			}
		}
	}
}
