package _095_unique_binary_search_trees_II

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

var empty = []*TreeNode{nil}

func generateTrees(n int) []*TreeNode {
	tries := [10][10][]*TreeNode{}
	for i := 0; i <= n; i++ {
		tries[i+1][i] = empty
	}

	for dist := 0; dist < n; dist++ {
		for start := 1; start <= n-dist; start++ {
			end := start + dist
			for root := start; root <= end; root++ {
				left, right := tries[start][root-1], tries[root+1][end]

				for _, l := range left {
					for _, r := range right {
						tries[start][end] = append(tries[start][end], &TreeNode{Val: root, Left: l, Right: r})
					}
				}
			}
		}
	}

	return tries[1][n]
}
