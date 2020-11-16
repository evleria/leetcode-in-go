package _110_balanced_binary_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func isBalanced(root *TreeNode) bool {
	_, ok := checkTree(root)

	return ok
}

func checkTree(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}

	leftDepth, leftOk := checkTree(node.Left)
	if !leftOk {
		return 0, false
	}

	rightDepth, rightOk := checkTree(node.Right)
	if !rightOk {
		return 0, false
	}

	if abs(leftDepth-rightDepth) > 1 {
		return 0, false
	}

	return max(leftDepth, rightDepth) + 1, true
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
