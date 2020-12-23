package _046_permutations

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestPermute(t *testing.T) {
	testCases := []struct {
		got  []int
		want [][]int
	}{
		{
			got: []int{1},
			want: [][]int{
				{1},
			},
		},
		{
			got: []int{0, 1},
			want: [][]int{
				{0, 1},
				{1, 0},
			},
		},
		{
			got: []int{1, 2, 3},
			want: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
	}

	for _, testCase := range testCases {
		actual := permute(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
