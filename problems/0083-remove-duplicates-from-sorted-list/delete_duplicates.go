package _083_remove_duplicates_from_sorted_list

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
)

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	current := head

	for current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}
