package _222_count_complete_tree_nodes

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)
	if leftHeight == rightHeight {
		leftCountWithRoot := 1 << leftHeight
		return leftCountWithRoot + countNodes(root.Right)
	}
	rightCountWithRoot := 1 << rightHeight
	return rightCountWithRoot + countNodes(root.Left)
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + getHeight(root.Left)
}
