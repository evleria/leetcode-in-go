package _827_making_a_large_island

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestLargestIsland(t *testing.T) {
	testCases := []struct {
		gotNums [][]int
		want    int
	}{
		{
			gotNums: [][]int{{0, 0}, {0, 0}},
			want:    1,
		},

		{
			gotNums: [][]int{{1, 0}, {0, 1}},
			want:    3,
		},
		{
			gotNums: [][]int{{1, 1}, {1, 0}},
			want:    4,
		},
		{
			gotNums: [][]int{{1, 1}, {1, 1}},
			want:    4,
		},
	}

	for _, testCase := range testCases {
		actual := largestIsland(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
