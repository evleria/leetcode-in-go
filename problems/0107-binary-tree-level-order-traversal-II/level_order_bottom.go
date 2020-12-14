package _107_binary_tree_level_order_traversal_II

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result, queue := make([][]int, 0, 16), make([]*TreeNode, 1, 16)
	queue[0] = root

	for len(queue) != 0 {
		level := make([]int, 0, 16)
		for _, cur := range queue {
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			level = append(level, cur.Val)
		}
		result = append([][]int{level}, result...)
	}

	return result
}
