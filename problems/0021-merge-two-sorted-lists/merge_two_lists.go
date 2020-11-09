package _021_merge_two_sorted_lists

import (
	. "github.com/evleria/leetcode-in-go/structs"
)

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := new(ListNode)
	current := dummyHead

	for l1 != nil && l2 != nil {
		var node *ListNode
		if l1.Val < l2.Val {
			node = l1
			l1 = l1.Next
		} else {
			node = l2
			l2 = l2.Next
		}
		current.Next = node
		current = node
	}

	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return dummyHead.Next
}
