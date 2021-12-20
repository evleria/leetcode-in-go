package _200_minimum_absolute_difference

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinimumAbsDifference(t *testing.T) {
	testCases := []struct {
		got  []int
		want [][]int
	}{
		{
			got:  []int{4, 2, 1, 3},
			want: [][]int{{1, 2}, {2, 3}, {3, 4}},
		},
		{
			got:  []int{1, 3, 6, 10, 15},
			want: [][]int{{1, 3}},
		},
		{
			got:  []int{3, 8, -10, 23, 19, -4, -14, 27},
			want: [][]int{{-14, -10}, {19, 23}, {23, 27}},
		},
	}

	for _, testCase := range testCases {
		actual := minimumAbsDifference(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
