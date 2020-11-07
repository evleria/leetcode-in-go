package _021_merge_two_sorted_lists

import (
	. "github.com/evleria/leetcode-in-go/structs"
)

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := new(ListNode)
	current := dummyHead

	for l1 != nil && l2 != nil {
		var val int
		if l1.Val < l2.Val {
			val = l1.Val
			l1 = l1.Next
		} else {
			val = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: val}
		current = current.Next
	}

	if l1 == nil {
		current.Next = l2
	} else {
		current.Next = l1
	}
	return dummyHead.Next
}
