package _111_minimum_depth_of_binary_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	min := minDepth(root.Left)
	if rightDepth := minDepth(root.Right); min == 0 || rightDepth < min && rightDepth != 0 {
		min = rightDepth
	}

	return min + 1
}
