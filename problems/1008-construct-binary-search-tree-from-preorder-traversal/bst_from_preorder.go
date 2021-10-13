package _008_construct_binary_search_tree_from_preorder_traversal

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) < 1 {
		return nil
	}
	bounds := 0
	for bounds < len(preorder) && preorder[bounds] <= preorder[0] {
		bounds++
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  bstFromPreorder(preorder[1:bounds]),
		Right: bstFromPreorder(preorder[bounds:]),
	}
}
