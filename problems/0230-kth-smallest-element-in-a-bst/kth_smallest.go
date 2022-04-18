package _230_kth_smallest_element_in_a_bst

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func kthSmallest(root *TreeNode, k int) int {
	var arr = []int{}
	n := arrayList(root, arr)
	return n[k-1]
}

func arrayList(root *TreeNode, arr []int) []int {
	if root == nil {
		return arr
	}
	arr = arrayList(root.Left, arr)
	arr = append(arr, root.Val)
	arr = arrayList(root.Right, arr)
	return arr
}
