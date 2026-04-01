package statistics_test

import (
	"fmt"
	"iter"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KrischanCS/go-toolbox/iterator"
	"github.com/KrischanCS/go-toolbox/iterator/reducer/statistics"
)

func ExampleMin() {
	i := iterator.Of(5, 2, 8, 1, 7)
	acc := math.MaxInt

	// Reduce a sequence to find the maximum value
	iterator.Reduce(i, &acc, statistics.Min[int])

	fmt.Println(acc)

	// Output: 1
}

func ExampleMax() {
	i := iterator.Of(5, 2, 8, 1, 7)
	acc := math.MinInt

	// Reduce a sequence to find the maximum value
	iterator.Reduce(i, &acc, statistics.Max[int])

	fmt.Println(acc)

	// Output: 8
}

func ExampleMean() {
	i := iterator.Of(0, -3, 5, 8, -2, 7)
	acc := statistics.MeanAccumulator[int]{}

	iterator.Reduce(i, &acc, statistics.Mean[int])

	fmt.Println(acc.Mean())
	// Output: 2.5
}

func ExampleMinMax() {
	i := iterator.Of(-6.28, 2.78, 9.81, 1.41)
	acc := statistics.NewMinMaxAccumulator[float64]()

	iterator.Reduce(i, &acc, statistics.MinMax[float64])

	fmt.Printf("Min: %.2f, Max: %.2f\n", acc.Min(), acc.Max())
	// Output: Min: -6.28, Max: 9.81
}

func TestMin(t *testing.T) {
	t.Parallel()

	type test struct {
		name         string
		input        iter.Seq[int]
		initialValue int
		expect       int
	}

	tests := []test{
		{
			name:         "Should not alter initial value if input is empty",
			input:        iterator.Of[int](),
			initialValue: math.MaxInt,
			expect:       math.MaxInt,
		},
		{
			name:         "Should return given value if only one is given",
			input:        iterator.Of[int](1),
			initialValue: math.MaxInt,
			expect:       1,
		},
		{
			name:         "Should return minimum value of input",
			input:        iterator.Of[int](1, 2, 3, 4, 5),
			initialValue: math.MaxInt,
			expect:       1,
		},
		{
			name:         "Should return minimum value of input with negative values",
			input:        iterator.Of[int](-1, -2, -3, -4, -5),
			initialValue: math.MaxInt,
			expect:       -5,
		},
		{
			name:         "Should return minimum value of input with mixed values",
			input:        iterator.Of[int](-1, 2, -3, 4, -5),
			initialValue: math.MaxInt,
			expect:       -5,
		},
		{
			name:         "Should work correctly with large numbers",
			input:        iterator.Of[int](math.MaxInt - 1),
			initialValue: math.MaxInt,
			expect:       math.MaxInt - 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			acc := tc.initialValue

			// Act
			iterator.Reduce(tc.input, &acc, statistics.Min[int])

			// Assert
			assert.Equal(t, tc.expect, acc)
		})
	}
}

func TestMax(t *testing.T) {
	t.Parallel()

	type test struct {
		name         string
		input        iter.Seq[int]
		initialValue int
		expect       int
	}

	tests := []test{
		{
			name:         "Should not alter initial value if input is empty",
			input:        iterator.Of[int](),
			initialValue: math.MinInt,
			expect:       math.MinInt,
		},
		{
			name:         "Should rerturn given value if only one is given",
			input:        iterator.Of[int](1),
			initialValue: math.MinInt,
			expect:       1,
		},
		{
			name:         "Should return maximum value of input",
			input:        iterator.Of[int](1, 2, 3, 4, 5),
			initialValue: math.MinInt,
			expect:       5,
		},
		{
			name:         "Should return maximum value of input with negative values",
			input:        iterator.Of[int](-1, -2, -3, -4, -5),
			initialValue: math.MinInt,
			expect:       -1,
		},
		{
			name:         "Should return maximum value of input with mixed values",
			input:        iterator.Of[int](-1, 2, -3, 4, -5),
			initialValue: math.MinInt,
			expect:       4,
		},
		{
			name:         "Should work correctly with near min numbers",
			input:        iterator.Of[int](math.MinInt + 1),
			initialValue: math.MinInt,
			expect:       math.MinInt + 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			acc := tc.initialValue

			// Act
			iterator.Reduce(tc.input, &acc, statistics.Max[int])

			// Assert
			assert.Equal(t, tc.expect, acc)
		})
	}
}

//nolint:funlen
func TestMean(t *testing.T) {
	t.Parallel()

	type test struct {
		name         string
		input        iter.Seq[int]
		initialValue statistics.MeanAccumulator[int]
		expect       float64
	}

	tests := []test{
		{
			name:         "Should return NaN if input is empty",
			input:        iterator.Of[int](),
			initialValue: statistics.MeanAccumulator[int]{},
			expect:       math.NaN(),
		},
		{
			name:         "Should return mean as the given value if only one is given",
			input:        iterator.Of[int](1),
			initialValue: statistics.MeanAccumulator[int]{},
			expect:       1,
		},
		{
			name:         "Should return mean correctly for two values",
			input:        iterator.Of[int](1, 2),
			initialValue: statistics.MeanAccumulator[int]{},
			expect:       1.5,
		},
		{
			name:         "Should return mean correctly for multiple values",
			input:        iterator.Of[int](1, 2, 3, 4, 5),
			initialValue: statistics.MeanAccumulator[int]{},
			expect:       3,
		},
		{
			name:         "Should return mean correctly with negative values",
			input:        iterator.Of[int](-1, -2, -3, -4, -5),
			initialValue: statistics.MeanAccumulator[int]{},
			expect:       -3,
		},
		{
			name:         "Should return mean correctly with mixed values",
			input:        iterator.Of[int](-1, 2, -3, 4, -5),
			initialValue: statistics.MeanAccumulator[int]{},
			expect:       -0.6,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			acc := tc.initialValue

			// Act
			iterator.Reduce(tc.input, &acc, statistics.Mean[int])

			// Assert
			if math.IsNaN(tc.expect) {
				assert.True(t, math.IsNaN(acc.Mean()))
				return
			}

			assert.InEpsilon(t, tc.expect, acc.Mean(), 0.0001)
		})
	}
}

//nolint:funlen
func TestMinMax(t *testing.T) {
	t.Parallel()

	type test struct {
		name         string
		input        iter.Seq[int]
		initialValue statistics.MinMaxAccumulator[int]
		expectMin    int
		expectMax    int
	}

	tests := []test{
		{
			name:         "Should not alter initial value if input is empty",
			input:        iterator.Of[int](),
			initialValue: statistics.NewMinMaxAccumulator[int](),
			expectMin:    math.MaxInt,
			expectMax:    math.MinInt,
		},
		{
			name:         "Should return min and max as the given value if only one is given",
			input:        iterator.Of[int](1),
			initialValue: statistics.NewMinMaxAccumulator[int](),
			expectMin:    1,
			expectMax:    1,
		},
		{
			name:         "Should return min and max correctly for two values",
			input:        iterator.Of[int](1, 2),
			initialValue: statistics.NewMinMaxAccumulator[int](),
			expectMin:    1,
			expectMax:    2,
		},
		{
			name:         "Should return min and max correctly for multiple values",
			input:        iterator.Of[int](1, 2, 3, 4, 5),
			initialValue: statistics.NewMinMaxAccumulator[int](),
			expectMin:    1,
			expectMax:    5,
		},
		{
			name:         "Should return min and max correctly with negative values",
			input:        iterator.Of[int](-1, -2, -3, -4, -5),
			initialValue: statistics.NewMinMaxAccumulator[int](),
			expectMin:    -5,
			expectMax:    -1,
		},
		{
			name:         "Should return min and max correctly with mixed values",
			input:        iterator.Of[int](-1, 2, -3, 4, -5),
			initialValue: statistics.NewMinMaxAccumulator[int](),
			expectMin:    -5,
			expectMax:    4,
		},
		{
			name:         "Should work correctly with large numbers",
			input:        iterator.Of[int](math.MaxInt-1, math.MinInt+1),
			initialValue: statistics.NewMinMaxAccumulator[int](),
			expectMin:    math.MinInt + 1,
			expectMax:    math.MaxInt - 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			acc := tc.initialValue

			// Act
			iterator.Reduce(tc.input, &acc, statistics.MinMax[int])

			// Assert
			assert.Equal(t, tc.expectMin, acc.Min())
			assert.Equal(t, tc.expectMax, acc.Max())
		})
	}
}
