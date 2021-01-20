package _059_spiral_matrix_II

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestGenerateMatrix(t *testing.T) {
	testCases := []struct {
		got  int
		want [][]int
	}{
		{
			got: 1,
			want: [][]int{
				{1},
			},
		},
		{
			got: 2,
			want: [][]int{
				{1, 2},
				{4, 3},
			},
		},
		{
			got: 3,
			want: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
	}

	for _, testCase := range testCases {
		actual := generateMatrix(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
