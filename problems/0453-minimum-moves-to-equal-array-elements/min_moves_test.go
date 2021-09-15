package _453_minimum_moves_to_equal_array_elements

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinMoves(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{-100, 0, 100},
			want: 300,
		},
		{
			got:  []int{1, 2, 3},
			want: 3,
		},
		{
			got:  []int{1, 1, 1},
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := minMoves(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
