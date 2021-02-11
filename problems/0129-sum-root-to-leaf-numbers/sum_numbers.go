package _129_sum_root_to_leaf_numbers

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, cur int) int {
	if root == nil {
		return 0
	}

	cur = cur*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return cur
	}
	return dfs(root.Left, cur) + dfs(root.Right, cur)
}
