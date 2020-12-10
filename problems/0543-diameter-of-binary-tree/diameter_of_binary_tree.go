package _543_diameter_of_binary_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result, _ := helper(root)
	return result
}

func helper(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	lDiameter, lDepth := helper(root.Left)
	rDiameter, rDepth := helper(root.Right)

	return max(lDepth+rDepth, max(lDiameter, rDiameter)), max(lDepth, rDepth) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
