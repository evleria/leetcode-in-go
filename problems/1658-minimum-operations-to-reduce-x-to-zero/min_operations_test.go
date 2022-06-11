package _658_minimum_operations_to_reduce_x_to_zero

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinOperations(t *testing.T) {
	testCases := []struct {
		gotNums []int
		gotX    int
		want    int
	}{
		{
			gotNums: []int{1, 1, 4, 2, 3},
			gotX:    5,
			want:    2,
		},
		{
			gotNums: []int{5, 6, 7, 8, 9},
			gotX:    4,
			want:    -1,
		},
		{
			gotNums: []int{3, 2, 20, 1, 1, 3},
			gotX:    10,
			want:    5,
		},
	}

	for _, testCase := range testCases {
		actual := minOperations(testCase.gotNums, testCase.gotX)

		assert.Check(t, is.Equal(actual, testCase.want))
	}
}
