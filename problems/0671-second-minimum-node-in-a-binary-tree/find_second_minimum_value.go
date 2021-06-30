package _671_second_minimum_node_in_a_binary_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func findSecondMinimumValue(root *TreeNode) int {
	return findMinNode(root, root.Val)
}

func findMinNode(root *TreeNode, except int) int {
	if root == nil {
		return -1
	}

	if root.Val != except {
		return root.Val
	}

	l := findMinNode(root.Left, except)
	if r := findMinNode(root.Right, except); l == -1 || (r != -1 && r < l) {
		return r
	}
	return l
}
