package _105_construct_binary_tree_from_preorder_and_inorder_traversal

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	idx := indexOf(inorder, preorder[0])
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:idx+1], inorder[:idx]),
		Right: buildTree(preorder[idx+1:], inorder[idx+1:]),
	}
}

func indexOf(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}
