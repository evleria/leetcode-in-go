package _102_binary_tree_level_order_traversal

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		var level []int
		for _, node := range queue {
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
	}

	return result
}
