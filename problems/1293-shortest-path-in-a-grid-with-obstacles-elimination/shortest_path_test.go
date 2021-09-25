package _293_shortest_path_in_a_grid_with_obstacles_elimination

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestShortestPath(t *testing.T) {
	testCases := []struct {
		got  [][]int
		gotK int
		want int
	}{
		{
			got: [][]int{
				{0, 0, 0},
				{1, 1, 0},
				{0, 0, 0},
				{0, 1, 1},
				{0, 0, 0}},
			gotK: 1,
			want: 6,
		},
		{
			got: [][]int{
				{0, 1, 1},
				{1, 1, 1},
				{1, 0, 0}},
			gotK: 1,
			want: -1,
		},
	}

	for _, testCase := range testCases {
		actual := shortestPath(testCase.got, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
