package _141_linked_list_cycle

import (
	. "github.com/evleria/leetcode-in-go/structs"
)

func hasCycle(head *ListNode) bool {
	tortoise, hare := head, head

	for hare != nil && hare.Next != nil {
		tortoise = tortoise.Next
		hare = hare.Next.Next
		if tortoise == hare {
			return true
		}
	}

	return false
}
