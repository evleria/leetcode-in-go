package _897_increasing_order_search_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func increasingBST(root *TreeNode) *TreeNode {
	dummy := new(TreeNode)
	prev := dummy
	traverse(root, &prev)
	return dummy.Right
}

func traverse(node *TreeNode, prev **TreeNode) {
	if node == nil {
		return
	}

	traverse(node.Left, prev)

	node.Left, (*prev).Right = nil, node
	*prev = node

	traverse(node.Right, prev)
}
