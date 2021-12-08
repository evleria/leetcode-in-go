package _563_binary_tree_tilt

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func findTilt(root *TreeNode) int {
	total := 0
	dfs(root, &total)
	return total
}

func dfs(root *TreeNode, total *int) int {
	if root == nil {
		return 0
	}

	left, right := dfs(root.Left, total), dfs(root.Right, total)
	*total += abs(left - right)
	return left + right + root.Val
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
