package _572_matrix_diagonal_sum

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestDiagonalSum(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want int
	}{
		{
			got: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9}},
			want: 25,
		},
		{
			got: [][]int{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1}},
			want: 8,
		},
		{
			got:  [][]int{{5}},
			want: 5,
		},
	}

	for _, testCase := range testCases {
		actual := diagonalSum(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
