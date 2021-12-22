package _143_reorder_list

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
)

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	cur := head
	for ; cur.Next.Next != nil; cur = cur.Next {
	}
	head.Next, cur.Next, cur.Next.Next = cur.Next, cur.Next.Next, head.Next
	reorderList(head.Next.Next)
}
