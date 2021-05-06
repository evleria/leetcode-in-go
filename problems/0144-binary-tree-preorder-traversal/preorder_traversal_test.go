package _144_binary_tree_preorder_traversal

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestPreorderTraversal(t *testing.T) {
	testCases := []struct {
		got  []int
		want []int
	}{
		{
			got:  []int{1, NULL, 2, 3},
			want: []int{1, 2, 3},
		},
	}

	for _, testCase := range testCases {
		actual := preorderTraversal(BinaryTreeFromSlice(testCase.got))

		assert.Check(t, is.DeepEqual(actual, testCase.want), testCase.got)
	}
}
