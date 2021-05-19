package _462_minimum_moves_to_equal_array_elements_II

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinMoves2(t *testing.T) {
	testCases := []struct {
		gotNums []int
		want    int
	}{
		{
			gotNums: []int{1, 2, 3},
			want:    2,
		},
		{
			gotNums: []int{1, 10, 2, 9},
			want:    16,
		},
	}

	for _, testCase := range testCases {
		actual := minMoves2(testCase.gotNums)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
