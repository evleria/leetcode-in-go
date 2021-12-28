package _876_middle_of_the_linked_list

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
)

func middleNode(head *ListNode) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Next = head
	tortoise, hare := dummyHead, dummyHead.Next
	for hare != nil && hare.Next != nil {
		tortoise = tortoise.Next
		hare = hare.Next.Next
	}
	return tortoise.Next
}
