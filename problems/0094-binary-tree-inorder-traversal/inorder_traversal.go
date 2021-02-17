package _094_binary_tree_inorder_traversal

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	inorderWrite(&result, root)
	return result
}

func inorderWrite(result *[]int, root *TreeNode) {
	if root != nil {
		inorderWrite(result, root.Left)
		*result = append(*result, root.Val)
		inorderWrite(result, root.Right)
	}
}
