package _106_construct_binary_tree_from_inorder_and_postorder_traversal

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	last := len(postorder) - 1
	idx := indexOf(inorder, postorder[last])
	return &TreeNode{
		Val:   postorder[last],
		Left:  buildTree(inorder[:idx], postorder[:idx]),
		Right: buildTree(inorder[idx+1:], postorder[idx:last]),
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
