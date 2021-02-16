package _144_binary_tree_preorder_traversal

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	preorderWrite(&result, root)
	return result
}

func preorderWrite(result *[]int, root *TreeNode) {
	if root != nil {
		*result = append(*result, root.Val)
		preorderWrite(result, root.Left)
		preorderWrite(result, root.Right)
	}
}
