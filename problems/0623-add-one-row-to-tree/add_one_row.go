package _623_add_one_row_to_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if root == nil {
		return nil
	}
	switch d {
	case 1:
		return &TreeNode{Val: v, Left: root}
	case 2:
		root.Left, root.Right = &TreeNode{Val: v, Left: root.Left}, &TreeNode{Val: v, Right: root.Right}
	default:
		addOneRow(root.Left, v, d-1)
		addOneRow(root.Right, v, d-1)
	}
	return root
}
