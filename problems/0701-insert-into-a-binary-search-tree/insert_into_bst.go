package _701_insert_into_a_binary_search_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}

	return root
}
