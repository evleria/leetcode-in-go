package _368_largest_divisible_subset

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestLargestDivisibleSubset(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    []int
	}{
		{
			gotNums: []int{1, 2, 3},
			want:    []int{1, 2},
		},
		{
			gotNums: []int{1, 2, 4, 8},
			want:    []int{1, 2, 4, 8},
		},
	}

	for _, testCase := range testCases {
		actual := largestDivisibleSubset(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
