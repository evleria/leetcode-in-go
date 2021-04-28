package _063_unique_paths_II

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want int
	}{
		{
			got: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			want: 2,
		},
		{
			got: [][]int{
				{0, 1},
				{0, 0},
			},
			want: 1,
		},
	}

	for _, testCase := range testCases {
		actual := uniquePathsWithObstacles(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
