package _365_how_many_numbers_are_smaller_than_the_current_number

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSmallerThanCurrent(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    []int
	}{
		{
			gotNums: []int{8, 1, 2, 2, 3},
			want:    []int{4, 0, 1, 1, 3},
		},
		{
			gotNums: []int{6, 5, 4, 8},
			want:    []int{2, 1, 0, 3},
		},
		{
			gotNums: []int{7, 7, 7, 7},
			want:    []int{0, 0, 0, 0},
		},
	}

	for _, testCase := range testCases {
		actual := smallerNumbersThanCurrent(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
