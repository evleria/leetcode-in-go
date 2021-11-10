package _748_sum_of_unique_elements

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSumOfUnique(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{1, 2, 3, 2},
			want:    4,
		},
		{
			gotNums: []int{1, 1, 1, 1, 1},
			want:    0,
		},
		{
			gotNums: []int{1, 2, 3, 4, 5},
			want:    15,
		},
	}

	for _, testCase := range testCases {
		actual := sumOfUnique(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
