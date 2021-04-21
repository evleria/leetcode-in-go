package _120_triangle

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinimumTotal(t *testing.T) {
	testCases := []struct {
		got  [][]int
		want int
	}{
		{
			got: [][]int{
				{2},
				{3, 4},
				{6, 5, 7},
				{4, 1, 8, 3},
			},
			want: 11,
		},
		{
			got: [][]int{
				{-10},
			},
			want: -10,
		},
	}

	for _, testCase := range testCases {
		actual := minimumTotal(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
