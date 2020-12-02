package _226_invert_binary_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

	return root
}
