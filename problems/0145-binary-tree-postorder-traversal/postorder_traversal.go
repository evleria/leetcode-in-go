package _145_binary_tree_postorder_traversal

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	postorderWrite(&result, root)
	return result
}

func postorderWrite(result *[]int, root *TreeNode) {
	if root != nil {
		postorderWrite(result, root.Left)
		postorderWrite(result, root.Right)
		*result = append(*result, root.Val)
	}
}
