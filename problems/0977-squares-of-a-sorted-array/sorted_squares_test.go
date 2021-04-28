package _977_squares_of_a_sorted_array

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSortedSquares(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    []int
	}{
		{
			gotNums: []int{-4, -1, 0, 3, 10},
			want:    []int{0, 1, 9, 16, 100},
		},
		{
			gotNums: []int{-7, -3, 2, 3, 11},
			want:    []int{4, 9, 9, 49, 121},
		},
	}

	for _, testCase := range testCases {
		actual := sortedSquares(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
