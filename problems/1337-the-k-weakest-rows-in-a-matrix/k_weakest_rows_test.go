package _337_the_k_weakest_rows_in_a_matrix

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestKWeakestRows(t *testing.T) {
	testCases := []struct {
		gotMatrix [][]int
		gotK      int
		want      []int
	}{
		{
			gotMatrix: [][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 1},
			},
			gotK: 3,
			want: []int{2, 0, 3},
		},
		{
			gotMatrix: [][]int{
				{1, 1},
				{1, 1},
				{1, 1},
			},
			gotK: 2,
			want: []int{0, 1},
		},
		{
			gotMatrix: [][]int{
				{1, 0},
				{0, 0},
				{1, 0},
			},
			gotK: 2,
			want: []int{1, 0},
		},
	}

	for _, testCase := range testCases {
		actual := kWeakestRows(testCase.gotMatrix, testCase.gotK)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
