package _112_path_sum

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}

	sum -= root.Val
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}
