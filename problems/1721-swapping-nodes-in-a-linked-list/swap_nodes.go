package _721_swapping_nodes_in_a_linked_list

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
)

func swapNodes(head *ListNode, k int) *ListNode {
	hi, lo := head, head
	for i := 1; i < k; i++ {
		hi = hi.Next
	}
	first := hi

	for hi.Next != nil {
		lo, hi = lo.Next, hi.Next
	}
	last := lo

	first.Val, last.Val = last.Val, first.Val

	return head
}
