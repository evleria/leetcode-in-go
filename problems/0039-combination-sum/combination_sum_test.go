package _039_combination_sum

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCombinationSum(t *testing.T) {
	testCases := []struct {
		got       []int
		gotTarget int
		want      [][]int
	}{
		{
			got:       []int{2, 3, 6, 7},
			gotTarget: 7,
			want: [][]int{
				{2, 2, 3},
				{7},
			},
		},
		{
			got:       []int{2, 3, 5},
			gotTarget: 8,
			want: [][]int{
				{2, 2, 2, 2},
				{2, 3, 3},
				{3, 5},
			},
		},
		{
			got:       []int{1},
			gotTarget: 2,
			want: [][]int{
				{1, 1},
			},
		},
	}

	for _, testCase := range testCases {
		actual := combinationSum(testCase.got, testCase.gotTarget)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
