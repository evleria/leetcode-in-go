package _718_maximum_length_of_repeated_subarray

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindLength(t *testing.T) {
	testCases := []struct {
		gotNums1 []int
		gotNums2 []int
		want     int
	}{
		{
			gotNums1: []int{1, 2, 3, 2, 1},
			gotNums2: []int{3, 2, 1, 4, 7},
			want:     3,
		},
		{
			gotNums1: []int{0, 0, 0, 0, 0},
			gotNums2: []int{0, 0, 0, 0, 0},
			want:     5,
		},
	}

	for _, testCase := range testCases {
		actual := findLength(testCase.gotNums1, testCase.gotNums2)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
