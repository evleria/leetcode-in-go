package _114_flatten_binary_tree_to_linked_list

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	stack := []*TreeNode{root.Right, root.Left}
	for len(stack) != 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if n != nil {
			stack = append(stack, n.Right, n.Left)
			root.Left, root.Right, root = nil, n, n
		}
	}
}
