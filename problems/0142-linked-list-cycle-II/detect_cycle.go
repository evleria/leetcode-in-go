package _142_linked_list_cycle_II

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
)

func detectCycle(head *ListNode) *ListNode {
	tortoise, hare, isCycle := head, head, false
	for hare != nil && hare.Next != nil {
		tortoise, hare = tortoise.Next, hare.Next.Next
		if tortoise == hare {
			isCycle = true
			break
		}
	}

	if !isCycle {
		return nil
	}

	tortoise = head
	for tortoise != hare {
		tortoise, hare = tortoise.Next, hare.Next
	}
	return tortoise
}
