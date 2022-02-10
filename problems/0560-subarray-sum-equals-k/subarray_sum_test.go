package _560_subarray_sum_equals_k

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSubarraySum(t *testing.T) {
	testCases := []struct {
		gotNums []int
		gotK    int
		want    int
	}{
		{
			gotNums: []int{1, 1, 1},
			gotK:    2,
			want:    2,
		},
		{
			gotNums: []int{1, 2, 3},
			gotK:    3,
			want:    2,
		},
	}

	for _, testCase := range testCases {
		actual := subarraySum(testCase.gotNums, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
