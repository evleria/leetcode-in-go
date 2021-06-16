package _404_sum_of_left_leaves

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum += root.Left.Val
	}
	return sum
}
