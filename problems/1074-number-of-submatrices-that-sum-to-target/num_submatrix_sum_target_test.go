package _074_number_of_submatrices_that_sum_to_target

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNumSubmatrixSumTarget(t *testing.T) {
	testCases := []struct {
		got  [][]int
		gotT int
		want int
	}{
		{
			got: [][]int{
				{0, 1, 0},
				{1, 1, 1},
				{0, 1, 0},
			},
			gotT: 0,
			want: 4,
		},
		{
			got: [][]int{
				{1, -1},
				{-1, 1},
			},
			gotT: 0,
			want: 5,
		},
		{
			got: [][]int{
				{904},
			},
			gotT: 0,
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := numSubmatrixSumTarget(testCase.got, testCase.gotT)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
