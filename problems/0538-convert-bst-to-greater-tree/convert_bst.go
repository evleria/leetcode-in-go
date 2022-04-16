package _538_convert_bst_to_greater_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func convertBST(root *TreeNode) *TreeNode {
	var sum int
	traverse(root, &sum)
	return root
}

func traverse(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	traverse(root.Right, sum)
	root.Val += *sum
	*sum = root.Val
	traverse(root.Left, sum)
}
