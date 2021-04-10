package _329_longest_increasing_path_in_a_matrix

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestLongestIncreasingPath(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want int
	}{
		{
			got: [][]int{
				{9, 9, 4},
				{6, 6, 8},
				{2, 1, 1},
			},
			want: 4,
		},
		{
			got: [][]int{
				{3, 4, 5},
				{3, 2, 6},
				{2, 2, 1},
			},
			want: 4,
		},
	}

	for _, testCase := range testCases {
		actual := longestIncreasingPath(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
