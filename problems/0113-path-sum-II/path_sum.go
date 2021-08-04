package _113_path_sum_II

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func pathSum(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}
	dfs(root, nil, targetSum, &result)
	return result
}

func dfs(node *TreeNode, path []int, leftSum int, result *[][]int) {
	if node == nil {
		return
	}

	leftSum -= node.Val
	newPath := append(path, node.Val)
	if node.Left == nil && node.Right == nil && leftSum == 0 {
		cp := make([]int, len(newPath))
		copy(cp, newPath)
		*result = append(*result, cp)
	} else {
		dfs(node.Left, newPath, leftSum, result)
		dfs(node.Right, newPath, leftSum, result)
	}
}
