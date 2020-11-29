package _053_maximum_subarray

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMaxSubArray(t *testing.T) {
	testCases := []struct {
		got  []int
		want int
	}{
		{
			got:  []int{-1},
			want: -1,
		},
		{
			got:  []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
	}

	for _, testCase := range testCases {
		actual := maxSubArray(testCase.got)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
