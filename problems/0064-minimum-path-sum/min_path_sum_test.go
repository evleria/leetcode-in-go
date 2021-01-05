package _064_minimum_path_sum

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinPathSum(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want int
	}{
		{
			got: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			want: 7,
		},
		{
			got: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			want: 12,
		},
	}

	for _, testCase := range testCases {
		actual := minPathSum(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
