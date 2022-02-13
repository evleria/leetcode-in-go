package _700_search_in_a_binary_search_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if val > root.Val {
		return searchBST(root.Right, val)
	}
	return searchBST(root.Left, val)
}
