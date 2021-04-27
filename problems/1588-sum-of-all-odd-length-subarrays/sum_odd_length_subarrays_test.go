package _588_sum_of_all_odd_length_subarrays

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSumOddLengthSubarrays(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{1, 4, 2, 5, 3},
			want:    58,
		},
		{
			gotNums: []int{1, 2},
			want:    3,
		},
		{
			gotNums: []int{10, 11, 12},
			want:    66,
		},
	}

	for _, testCase := range testCases {
		actual := sumOddLengthSubarrays(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
