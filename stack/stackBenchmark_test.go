package stack_test

import (
	"slices"
	"testing"

	"github.com/KrischanCS/go-toolbox/iterator"
	"github.com/KrischanCS/go-toolbox/stack"
)

func BenchmarkStack_PeekN(b *testing.B) {
	b.ReportAllocs()

	s := stack.Of(
		slices.Collect(iterator.FromTo(0, 10_000))...)

	for b.Loop() {
		_ = s.PeekN(500)
	}
}
