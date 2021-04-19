package _377_combination_sum_IV

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestCombinationSum4(t *testing.T) {
	testCases := []struct {
		gotNums   []int
		gotTarget int
		want      int
	}{
		{
			gotNums:   []int{1, 2, 3},
			gotTarget: 4,
			want:      7,
		},
		{
			gotNums:   []int{9},
			gotTarget: 3,
			want:      0,
		},
		{
			gotNums:   []int{3, 1, 2, 4},
			gotTarget: 4,
			want:      8,
		},
	}

	for _, testCase := range testCases {
		actual := combinationSum4(testCase.gotNums, testCase.gotTarget)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
