package _310_minimum_height_trees

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindMinHeightTrees(t *testing.T) {
	testCases := []struct {
		gotN     int
		gotEdges [][]int
		want     []int
	}{
		{
			gotN:     4,
			gotEdges: [][]int{{1, 0}, {1, 2}, {1, 3}},
			want:     []int{1},
		},
		{
			gotN:     6,
			gotEdges: [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}},
			want:     []int{3, 4},
		},
		{
			gotN:     1,
			gotEdges: [][]int{},
			want:     []int{0},
		},
	}

	for _, testCase := range testCases {
		actual := findMinHeightTrees(testCase.gotN, testCase.gotEdges)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
