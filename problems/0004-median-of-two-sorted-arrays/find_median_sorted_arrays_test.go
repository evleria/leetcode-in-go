package _004_median_of_two_sorted_arrays

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindMedianSortedArrays(t *testing.T) {
	testCases := []struct {
		gotNums1 []int
		gotNums2 []int
		want     float64
	}{
		{
			gotNums1: []int{1, 3},
			gotNums2: []int{2},
			want:     2,
		},
		{
			gotNums1: []int{1, 2},
			gotNums2: []int{3, 4},
			want:     2.5,
		},
	}

	for _, testCase := range testCases {
		actual := findMedianSortedArrays(testCase.gotNums1, testCase.gotNums2)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
