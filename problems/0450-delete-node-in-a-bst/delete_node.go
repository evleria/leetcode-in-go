package _450_delete_node_in_a_bst

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			successor := root.Right
			for successor.Left != nil {
				successor = successor.Left
			}
			root.Val = successor.Val
			root.Right = deleteNode(root.Right, successor.Val)
			return root
		}
	}
}
