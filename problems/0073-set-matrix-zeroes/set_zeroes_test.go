package _073_set_matrix_zeroes

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSetZeroes(t *testing.T) {
	testCases := []struct {
		gotNums [][]int
		want    [][]int
	}{
		{
			gotNums: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			want:    [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			gotNums: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			want:    [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
	}

	for _, testCase := range testCases {
		setZeroes(testCase.gotNums)

		assert.Check(t, is.DeepEqual(testCase.gotNums, testCase.want))
	}
}
