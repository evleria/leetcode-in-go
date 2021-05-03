package _480_running_sum_of_1d_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestRunningSum(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    []int
	}{
		{
			gotNums: []int{1, 2, 3, 4},
			want:    []int{1, 3, 6, 10},
		},
		{
			gotNums: []int{1, 1, 1, 1, 1},
			want:    []int{1, 2, 3, 4, 5},
		},
	}

	for _, testCase := range testCases {
		actual := runningSum(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
