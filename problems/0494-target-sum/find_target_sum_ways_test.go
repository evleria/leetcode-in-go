package _494_target_sum

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindTargetSumWays(t *testing.T) {
	testCases := []struct {
		gotNums   []int
		gotTarget int
		want      int
	}{
		{
			gotNums:   []int{1, 1, 1, 1, 1},
			gotTarget: 3,
			want:      5,
		},
		{
			gotNums:   []int{1, 0},
			gotTarget: 1,
			want:      2,
		},
		{
			gotNums:   []int{0, 0, 1},
			gotTarget: 1,
			want:      4,
		},
	}

	for _, testCase := range testCases {
		actual := findTargetSumWays(testCase.gotNums, testCase.gotTarget)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
