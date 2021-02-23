package _337_house_robber_III

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func rob(root *TreeNode) int {
	return max(helper(root))
}

func helper(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}
	gold := node.Val
	l1, l2 := helper(node.Left)
	r1, r2 := helper(node.Right)
	c1 := gold + l2 + r2
	c2 := max(l1, l2) + max(r1, r2)
	return c1, c2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
