package _004_max_consecutive_ones_III

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestLongestOnes(t *testing.T) {
	testCases := []struct {
		gotNums []int
		gotK    int
		want    int
	}{
		{
			gotNums: []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			gotK:    2,
			want:    6,
		},
		{
			gotNums: []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			gotK:    3,
			want:    10,
		},
		{
			gotNums: []int{0, 0, 0},
			gotK:    0,
			want:    0,
		},
	}

	for _, testCase := range testCases {
		actual := longestOnes(testCase.gotNums, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
