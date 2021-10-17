package _437_path_sum_III

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func pathSum(root *TreeNode, targetSum int) int {
	return helper(root, targetSum, false)
}

func helper(root *TreeNode, targetSum int, started bool) int {
	if root == nil {
		return 0
	}

	sum := 0
	if targetSum == root.Val {
		sum++
	}
	sum += helper(root.Left, targetSum-root.Val, true) + helper(root.Right, targetSum-root.Val, true)
	if !started {
		sum += helper(root.Left, targetSum, false) + helper(root.Right, targetSum, false)
	}

	return sum
}
