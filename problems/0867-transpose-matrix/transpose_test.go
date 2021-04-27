package _867_transpose_matrix

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestTranspose(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want [][]int
	}{
		{
			got: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			},
		},
		{
			got: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			want: [][]int{
				{1, 4},
				{2, 5},
				{3, 6},
			},
		},
	}

	for _, testCase := range testCases {
		actual := transpose(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
