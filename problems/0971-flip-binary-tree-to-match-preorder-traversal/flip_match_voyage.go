package _971_flip_binary_tree_to_match_preorder_traversal

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	stack := []*TreeNode{root}
	result, i := []int{}, 0
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if voyage[i] != node.Val {
			return []int{-1}
		}
		i++
		if node.Right != nil && node.Left != nil && node.Left.Val != voyage[i] {
			node.Left, node.Right, result = node.Right, node.Left, append(result, node.Val)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}
