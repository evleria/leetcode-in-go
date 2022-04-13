package _236_root_equals_sum_of_children

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func checkTree(root *TreeNode) bool {
	if root.Left.Val+root.Right.Val == root.Val {
		return true
	}
	return false
}
