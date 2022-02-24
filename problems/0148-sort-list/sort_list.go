package _148_sort_list

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
)

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// Split
	first, second := head, head
	for second.Next != nil && second.Next.Next != nil {
		first, second = first.Next, second.Next.Next
	}
	second = sortList(first.Next)
	first.Next = nil
	first = sortList(head)

	// Merge
	dh := new(ListNode)
	cur := dh

	for ; first != nil && second != nil; cur = cur.Next {
		if first.Val < second.Val {
			cur.Next, first = first, first.Next
		} else {
			cur.Next, second = second, second.Next
		}
	}

	if first != nil {
		cur.Next = first
	} else {
		cur.Next = second
	}

	return dh.Next
}
