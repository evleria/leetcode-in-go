package _695_max_area_of_island

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaxAreaOfIsland(t *testing.T) {
	testCases := []struct {
		gotGrid [][]int
		want    int
	}{
		{
			gotGrid: [][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 0, 1, 1},
				{0, 0, 0, 1, 1},
			},
			want: 4,
		},
		{
			gotGrid: [][]int{
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			},
			want: 6,
		},
		{
			gotGrid: [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := maxAreaOfIsland(testCase.gotGrid)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
