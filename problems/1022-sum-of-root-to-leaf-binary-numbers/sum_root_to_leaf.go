package _022_sum_of_root_to_leaf_binary_numbers

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func sumRootToLeaf(root *TreeNode) int {
	sum := 0
	dfs(root, 0, &sum)
	return sum
}

func dfs(root *TreeNode, base int, sum *int) {
	if root == nil {
		return
	}
	base = base<<1 + root.Val
	if root.Left == nil && root.Right == nil {
		*sum += base
	} else {
		dfs(root.Left, base, sum)
		dfs(root.Right, base, sum)
	}
}
