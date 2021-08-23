package _653_two_sum_IV_input_is_a_bst

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func findTarget(root *TreeNode, k int) bool {
	return dfs(root, k, map[int]bool{})
}

func dfs(root *TreeNode, k int, m map[int]bool) bool {
	if root == nil {
		return false
	}

	if m[k-root.Val] {
		return true
	}
	m[root.Val] = true

	return dfs(root.Left, k, m) || dfs(root.Right, k, m)
}
