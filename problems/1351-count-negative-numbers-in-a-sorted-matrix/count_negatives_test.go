package _351_count_negative_numbers_in_a_sorted_matrix

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCountNegative(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want int
	}{
		{
			got: [][]int{
				{4, 3, 2, -1},
				{3, 2, 1, -1},
				{1, 1, -1, -2},
				{-1, -1, -2, -3},
			},
			want: 8,
		},
		{
			got: [][]int{
				{3, 2},
				{1, 0},
			},
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := countNegatives(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
