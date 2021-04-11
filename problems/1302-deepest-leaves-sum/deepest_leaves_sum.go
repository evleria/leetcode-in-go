package _302_deepest_leaves_sum

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func deepestLeavesSum(root *TreeNode) int {
	queue := []*TreeNode{root}
	lastSum := 0
	for len(queue) > 0 {
		levelSum := 0
		for _, node := range queue {
			queue = queue[1:]
			levelSum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		lastSum = levelSum
	}
	return lastSum
}
