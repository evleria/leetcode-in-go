package _206_reverse_linked_list

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
)

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, head, prev = prev, head.Next, head
	}

	return prev
}
