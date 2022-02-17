package _530_minimum_absolute_difference_in_bst

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"math"
)

func getMinimumDifference(root *TreeNode) int {
	minDiff, prev := math.MaxInt32, math.MinInt32
	traverse(root, &minDiff, &prev)

	return minDiff
}

func traverse(root *TreeNode, minDiff *int, prev *int) {
	if root == nil {
		return
	}
	traverse(root.Left, minDiff, prev)
	diff := root.Val - *prev
	if diff < *minDiff {
		*minDiff = diff
	}
	*prev = root.Val
	traverse(root.Right, minDiff, prev)
}
