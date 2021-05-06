package _102_binary_tree_level_order_traversal

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestLevelOrder(t *testing.T) {
	testCases := []struct {
		got  []int
		want [][]int
	}{
		{
			got: []int{3, 9, 20, NULL, NULL, 15, 7},
			want: [][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
	}

	for _, testCase := range testCases {
		actual := levelOrder(BinaryTreeFromSlice(testCase.got))

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
