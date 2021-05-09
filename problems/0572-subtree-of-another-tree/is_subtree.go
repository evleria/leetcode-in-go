package _572_subtree_of_another_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func isSubtree(root, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return isEqual(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isEqual(a, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}
	return a.Val == b.Val && isEqual(a.Left, b.Left) && isEqual(a.Right, b.Right)
}
