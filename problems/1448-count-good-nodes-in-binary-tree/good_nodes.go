package _448_count_good_nodes_in_binary_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"math"
)

func goodNodes(root *TreeNode) int {
	sum := 0
	helper(root, math.MinInt64, &sum)
	return sum
}

func helper(root *TreeNode, m int, sum *int) {
	if root == nil {
		return
	}

	if root.Val >= m {
		*sum++
		m = root.Val
	}

	helper(root.Left, m, sum)
	helper(root.Right, m, sum)
}
