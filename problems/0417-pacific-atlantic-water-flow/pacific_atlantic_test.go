package _417_pacific_atlantic_water_flow

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestPacificAtlantic(t *testing.T) {
	testCases := []struct {
		gotNums [][]int
		want    [][]int
	}{
		{
			gotNums: [][]int{
				{3, 3, 3},
				{3, 1, 3},
				{0, 2, 4},
			},
			want: [][]int{
				{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}, {2, 2},
			},
		},
		{
			gotNums: [][]int{},
			want:    [][]int{},
		},
		{
			gotNums: [][]int{{1, 1}, {1, 1}, {1, 1}},
			want: [][]int{
				{0, 0}, {0, 1}, {1, 0}, {1, 1}, {2, 0}, {2, 1},
			},
		},
		{
			gotNums: [][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			},
			want: [][]int{
				{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0},
			},
		},
	}

	for _, testCase := range testCases {
		actual := pacificAtlantic(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
