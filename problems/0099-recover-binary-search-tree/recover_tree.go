package _099_recover_binary_search_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"math"
)

func recoverTree(root *TreeNode) {
	var first, second *TreeNode
	prev := &TreeNode{Val: math.MinInt64}

	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		traverse(node.Left)

		if first == nil && prev.Val >= node.Val {
			first = prev
		}
		if first != nil && prev.Val >= node.Val {
			second = node
		}

		prev = node
		traverse(node.Right)
	}

	traverse(root)
	first.Val, second.Val = second.Val, first.Val
}
