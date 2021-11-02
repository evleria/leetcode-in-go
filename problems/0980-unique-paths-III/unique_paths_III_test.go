package _980_unique_paths_III

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestUniquePathsIII(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want int
	}{
		{
			got: [][]int{
				{1, 0},
				{2, 0},
			},
			want: 1,
		},
		{
			got: [][]int{
				{1, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 2, -1},
			},
			want: 2,
		},
		{
			got: [][]int{
				{1, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 2},
			},
			want: 4,
		},
		{
			got: [][]int{
				{0, 1},
				{2, 0},
			},
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := uniquePathsIII(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
