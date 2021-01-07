package _018_4Sum

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFourSum(t *testing.T) {
	testCases := []struct {
		got    []int
		target int
		want   [][]int
	}{
		{
			got:    []int{},
			target: 0,
			want:   [][]int{},
		},
		{
			got:    []int{1, 0, -1, 0, -2, 2},
			target: 0,
			want: [][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
	}

	for _, testCase := range testCases {
		actual := fourSum(testCase.got, testCase.target)

		assert.Check(t, is.DeepEqual(actual, testCase.want))
	}
}
