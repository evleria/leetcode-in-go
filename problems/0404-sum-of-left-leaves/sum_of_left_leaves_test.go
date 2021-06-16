package _404_sum_of_left_leaves

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSumOfLeftLeaves(t *testing.T) {
	testCases := []struct {
		gotRoot []int
		want    int
	}{
		{
			gotRoot: []int{1, 2, 3, 4, 5},
			want:    4,
		},
		{
			gotRoot: []int{3, 9, 20, NULL, NULL, 15, 7},
			want:    24,
		},
		{
			gotRoot: []int{1},
			want:    0,
		},
	}

	for _, testCase := range testCases {
		actual := sumOfLeftLeaves(BinaryTreeFromSlice(testCase.gotRoot))

		assert.Check(t, is.Equal(actual, testCase.want), testCase.gotRoot)
	}
}
