package _153_find_minimum_in_rotated_sorted_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFindMin(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{3, 4, 5, 1, 2},
			want:    1,
		},
		{
			gotNums: []int{4, 5, 6, 7, 0, 1, 2},
			want:    0,
		},
		{
			gotNums: []int{11, 13, 15, 17},
			want:    11,
		},
	}

	for _, testCase := range testCases {
		actual := findMin(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
