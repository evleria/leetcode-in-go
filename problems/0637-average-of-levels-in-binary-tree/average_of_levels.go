package _637_average_of_levels_in_binary_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func averageOfLevels(root *TreeNode) []float64 {
	result, queue := []float64{}, []*TreeNode{root}
	for len(queue) > 0 {
		sum, count := 0, 0
		for _, node := range queue {
			queue = queue[1:]
			sum += node.Val
			count++
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, float64(sum)/float64(count))
	}

	return result
}
