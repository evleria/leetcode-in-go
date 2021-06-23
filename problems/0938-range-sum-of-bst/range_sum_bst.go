package _938_range_sum_of_bst

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	sum := rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
	if root.Val >= low && root.Val <= high {
		sum += root.Val
	}
	return sum
}
