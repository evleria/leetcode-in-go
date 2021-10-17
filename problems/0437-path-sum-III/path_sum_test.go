package _437_path_sum_III

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestPathSum(t *testing.T) {
	testCases := []struct {
		got       []int
		gotTarget int
		want      int
	}{
		{
			got:       []int{10, 5, -3, 3, 2, NULL, 11, 3, -2, NULL, 1},
			gotTarget: 8,
			want:      3,
		},
		{
			got:       []int{5, 4, 8, 11, NULL, 13, 4, 7, 2, NULL, NULL, 5, 1},
			gotTarget: 22,
			want:      3,
		},
	}

	for _, testCase := range testCases {
		actual := pathSum(BinaryTreeFromSlice(testCase.got), testCase.gotTarget)

		assert.Check(t, is.Equal(actual, testCase.want), testCase.got)
	}
}
