package _339_maximum_product_of_splitted_binary_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"math"
)

func maxProduct(root *TreeNode) int {
	s := sum(root)
	max := math.MinInt64
	dfs(root, s, &max)
	return max % 1_000_000_007
}

func dfs(root *TreeNode, total int, max *int) int {
	if root == nil {
		return 0
	}

	s := root.Val + dfs(root.Left, total, max) + dfs(root.Right, total, max)
	if r := s * (total - s); r > *max {
		*max = r
	}
	return s
}

func sum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return root.Val + sum(root.Left) + sum(root.Right)
}
