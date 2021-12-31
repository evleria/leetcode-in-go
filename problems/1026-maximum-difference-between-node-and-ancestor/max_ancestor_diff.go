package _026_maximum_difference_between_node_and_ancestor

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func maxAncestorDiff(root *TreeNode) int {
	return helper(root, root.Val, root.Val)
}

func helper(root *TreeNode, s, l int) int {
	if root == nil {
		return l - s
	}
	s = min(root.Val, s)
	l = max(root.Val, l)

	return max(helper(root.Left, s, l), helper(root.Right, s, l))
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
