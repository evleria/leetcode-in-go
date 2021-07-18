package _025_reverse_nodes_in_k_group

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
)

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	dummyHead := &ListNode{Next: head}
	cur := dummyHead

	for cur != nil {
		// go to the end of k-group. end will be the last element
		end := cur
		for i := 0; i < k; i++ {
			end = end.Next
			if end == nil {
				return dummyHead.Next
			}
		}

		// reverse nodes in between
		prev, next := cur, end.Next
		for x := cur.Next; x != next; {
			prev, x, x.Next = x, x.Next, prev
		}

		// link groups together
		cur, cur.Next, cur.Next.Next = cur.Next, end, next
	}

	return dummyHead.Next
}
