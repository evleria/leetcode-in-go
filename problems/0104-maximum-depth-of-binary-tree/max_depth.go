package _104_maximum_depth_of_binary_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := maxDepth(root.Left)
	if rightDepth := maxDepth(root.Right); rightDepth > max {
		max = rightDepth
	}

	return max + 1
}
