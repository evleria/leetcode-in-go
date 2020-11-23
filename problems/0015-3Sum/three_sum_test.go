package _015_3Sum

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestThreeSum(t *testing.T) {
	testCases := []struct {
		got  []int
		want [][]int
	}{
		{
			got:  []int{},
			want: [][]int{},
		},
		{
			got:  []int{0},
			want: [][]int{},
		},
		{
			got: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			got: []int{0, 1, -1, 2, 1, -1, 0},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
	}

	for _, testCase := range testCases {
		actual := threeSum(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
