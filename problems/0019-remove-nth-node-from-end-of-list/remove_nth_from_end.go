package _019_remove_nth_node_from_end_of_list

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	lo, hi := dummyHead, dummyHead

	for i := 0; i < n; i++ {
		hi = hi.Next
	}

	for hi.Next != nil {
		lo = lo.Next
		hi = hi.Next
	}

	lo.Next = lo.Next.Next

	return dummyHead.Next
}
