package _109_convert_sorted_list_to_binary_search_tree

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	. "github.com/evleria/leetcode-in-go/structs/tree"
)

func sortedListToBST(head *ListNode) *TreeNode {
	return helper(head, nil)
}

func helper(start, end *ListNode) *TreeNode {
	if start == end {
		return nil
	}

	if start.Next == end {
		return &TreeNode{Val: start.Val}
	}

	slow, fast := start, start

	for fast != end && fast.Next != end {
		slow, fast = slow.Next, fast.Next.Next
	}

	return &TreeNode{
		Val:   slow.Val,
		Left:  helper(start, slow),
		Right: helper(slow.Next, end),
	}
}
