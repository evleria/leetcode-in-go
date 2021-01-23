package _061_rotate_list

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
)

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}

	hi, lo, i := head, head, 0
	for ; i < k && hi != nil; i++ {
		hi = hi.Next
	}
	if hi == nil {
		return rotateRight(head, k%i)
	}
	for hi.Next != nil {
		hi, lo = hi.Next, lo.Next
	}
	head, lo.Next, hi.Next = lo.Next, nil, head
	return head
}
