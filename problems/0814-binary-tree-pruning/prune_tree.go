package _814_binary_tree_pruning

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = pruneTree(root.Left), pruneTree(root.Right)
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}
	return root
}
