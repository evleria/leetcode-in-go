package _001_two_sum

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		gotNums   []int
		gotTarget int
		want      []int
	}{
		{
			gotNums:   []int{2, 3},
			gotTarget: 5,
			want:      []int{0, 1},
		},
		{
			gotNums:   []int{2, 7, 11, 15},
			gotTarget: 9,
			want:      []int{0, 1},
		},
		{
			gotNums:   []int{3, 2, 4},
			gotTarget: 6,
			want:      []int{1, 2},
		},
		{
			gotNums:   []int{3, 3},
			gotTarget: 6,
			want:      []int{0, 1},
		},
	}

	for _, testCase := range testCases {
		actual := twoSum(testCase.gotNums, testCase.gotTarget)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
