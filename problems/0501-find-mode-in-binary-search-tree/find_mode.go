package _501_find_mode_in_binary_search_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"math"
)

func findMode(root *TreeNode) []int {
	result, count, localCount, prev := []int{}, 0, 0, math.MinInt64
	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Left)

		if prev == node.Val {
			localCount++
		} else {
			localCount = 1
		}

		switch {
		case localCount == count:
			result = append(result, node.Val)
		case localCount > count:
			count = localCount
			result = []int{node.Val}
		}

		prev = node.Val

		dfs(node.Right)
	}

	dfs(root)
	return result
}
