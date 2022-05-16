package _091_shortest_path_in_binary_matrix

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestShortestPathBinaryMatrix(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want int
	}{
		{
			got: [][]int{
				{0, 0, 0},
				{1, 1, 0},
				{1, 1, 0},
			},
			want: 4,
		},
		{
			got: [][]int{
				{0, 0, 0},
				{1, 1, 0},
				{1, 1, 0}},
			want: 4,
		},
		{
			got: [][]int{
				{1, 0, 0},
				{1, 1, 0},
				{1, 1, 0}},
			want: -1,
		},
		{
			got: [][]int{
				{0, 0, 0, 0, 1},
				{1, 0, 0, 0, 0},
				{0, 1, 0, 1, 0},
				{0, 0, 0, 1, 1},
				{0, 0, 0, 1, 0},
			},
			want: -1,
		},
	}

	for _, testCase := range testCases {
		actual := shortestPathBinaryMatrix(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
