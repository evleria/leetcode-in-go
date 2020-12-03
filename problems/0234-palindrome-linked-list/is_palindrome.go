package _234_palindrome_linked_list

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
)

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	var prev *ListNode
	for slow != nil {
		prev, slow, slow.Next = slow, slow.Next, prev
	}

	for prev != nil {
		if prev.Val != head.Val {
			return false
		}
		prev, head = prev.Next, head.Next
	}

	return true
}
