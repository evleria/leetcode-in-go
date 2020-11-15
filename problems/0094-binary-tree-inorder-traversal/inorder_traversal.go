package _094_binary_tree_inorder_traversal

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0, 4)

	writeInorder(root, &result)

	return result
}

func writeInorder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	writeInorder(node.Left, result)
	*result = append(*result, node.Val)
	writeInorder(node.Right, result)
}
