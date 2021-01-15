package _054_spiral_matrix

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSpiralOrder(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want []int
	}{
		{
			got: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			got: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			got: [][]int{
				{1, 2},
				{3, 4},
			},
			want: []int{1, 2, 4, 3},
		},
		{
			got: [][]int{
				{1},
			},
			want: []int{1},
		},
		{
			got: [][]int{
				{1},
				{2},
				{3},
			},
			want: []int{1, 2, 3},
		},
	}

	for _, testCase := range testCases {
		actual := spiralOrder(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
