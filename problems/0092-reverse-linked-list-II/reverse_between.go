package _092_reverse_linked_list_II

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
)

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	dummyHead := &ListNode{Next: head}
	preLeft := dummyHead
	for i := 0; i < left-1; i++ {
		head, preLeft = head.Next, head
	}

	var prev *ListNode
	preRight := head
	for i := 0; i < right-left+1; i++ {
		head, head.Next, prev = head.Next, prev, head
	}

	preLeft.Next = prev
	preRight.Next = head

	return dummyHead.Next
}
