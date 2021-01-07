package _040_combination_sum_II

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCombinationSum2(t *testing.T) {
	testCases := []struct {
		got       []int
		gotTarget int
		want      [][]int
	}{
		{
			got:       []int{10, 1, 2, 7, 6, 1, 5},
			gotTarget: 8,
			want: [][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
		},
		{
			got:       []int{2, 5, 2, 1, 2},
			gotTarget: 5,
			want: [][]int{
				{1, 2, 2},
				{5},
			},
		},
	}

	for _, testCase := range testCases {
		actual := combinationSum2(testCase.got, testCase.gotTarget)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
