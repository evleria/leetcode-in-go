package _203_remove_linked_list_elements

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
)

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}
	head = dummyHead

	for head.Next != nil {
		if head.Next.Val != val {
			head = head.Next
		} else {
			head.Next = head.Next.Next
		}
	}

	return dummyHead.Next
}
