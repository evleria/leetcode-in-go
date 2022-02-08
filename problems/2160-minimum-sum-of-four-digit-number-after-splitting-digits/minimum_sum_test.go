package _160_minimum_sum_of_four_digit_number_after_splitting_digits

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinimumSum(t *testing.T) {
	testCases := []struct {
		gotNums int
		want    int
	}{
		{
			gotNums: 2932,
			want:    52,
		},
		{
			gotNums: 4009,
			want:    13,
		},
	}

	for _, testCase := range testCases {
		actual := minimumSum(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
