package _235_lowest_common_ancestor_of_a_binary_search_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	switch {
	case p.Val < root.Val && q.Val < root.Val:
		return lowestCommonAncestor(root.Left, p, q)
	case p.Val > root.Val && q.Val > root.Val:
		return lowestCommonAncestor(root.Right, p, q)
	default:
		return root
	}
}
