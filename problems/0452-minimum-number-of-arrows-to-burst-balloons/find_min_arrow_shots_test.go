package _452_minimum_number_of_arrows_to_burst_balloons

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindMinArrowShots(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want int
	}{
		{
			got:  [][]int{{3, 9}, {7, 12}, {3, 8}, {6, 8}, {9, 10}, {2, 9}, {0, 9}, {3, 9}, {0, 6}, {2, 8}},
			want: 2,
		},
		{
			got:  [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}},
			want: 2,
		},
		{
			got:  [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			want: 4,
		},
		{
			got:  [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
			want: 2,
		},
	}

	for _, testCase := range testCases {
		actual := findMinArrowShots(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
