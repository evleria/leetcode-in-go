package _024_swap_nodes_in_pairs

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
)

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	head = dummyHead

	for head.Next != nil && head.Next.Next != nil {
		head, head.Next, head.Next.Next, head.Next.Next.Next =
			head.Next, head.Next.Next, head.Next.Next.Next, head.Next
	}

	return dummyHead.Next
}
