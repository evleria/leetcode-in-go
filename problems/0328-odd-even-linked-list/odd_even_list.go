package _328_odd_even_linked_list

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
)

func oddEvenList(head *ListNode) *ListNode {
	oddDh, evenDh := new(ListNode), new(ListNode)
	odd, even, isOdd := oddDh, evenDh, true
	for cur := head; cur != nil; cur, isOdd = cur.Next, !isOdd {
		if isOdd {
			odd.Next, odd = cur, cur
		} else {
			even.Next, even = cur, cur
		}
	}
	odd.Next, even.Next = evenDh.Next, nil
	return oddDh.Next
}
