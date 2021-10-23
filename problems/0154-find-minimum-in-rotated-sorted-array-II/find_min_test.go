package _154_find_minimum_in_rotated_sorted_array_II

import (
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestFindMin(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{1, 3, 5},
			want:    1,
		},
		{
			gotNums: []int{2, 2, 2, 0, 1},
			want:    0,
		},
	}

	for _, testCase := range testCases {
		actual := findMin(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
