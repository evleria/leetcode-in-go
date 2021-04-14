package _086_partition_list

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
)

func partition(head *ListNode, x int) *ListNode {
	lowerDummy, upperDummy := new(ListNode), &ListNode{Next: head}
	lowerEnd, current := lowerDummy, upperDummy
	for current.Next != nil {
		if current.Next.Val < x {
			current.Next, lowerEnd, lowerEnd.Next, current.Next.Next =
				current.Next.Next, current.Next, current.Next, lowerEnd.Next
		} else {
			current = current.Next
		}
	}
	lowerEnd.Next = upperDummy.Next
	return lowerDummy.Next
}
