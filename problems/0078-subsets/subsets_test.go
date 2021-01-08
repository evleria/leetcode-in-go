package _078_subsets

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSubsets(t *testing.T) {
	testCases := []struct {
		got  []int
		want [][]int
	}{
		{
			got: []int{1, 2, 3},
			want: [][]int{
				{},
				{1},
				{1, 2},
				{1, 2, 3},
				{1, 3},
				{2},
				{2, 3},
				{3},
			},
		},
		{
			got: []int{0},
			want: [][]int{
				{},
				{0},
			},
		},
	}

	for _, testCase := range testCases {
		actual := subsets(testCase.got)

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
