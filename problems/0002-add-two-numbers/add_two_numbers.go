package _002_add_two_numbers

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := new(ListNode)
	current := dummyHead
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := getValue(l1) + getValue(l2) + carry
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
		l1, l2 = getNext(l1), getNext(l2)
	}

	return dummyHead.Next
}

func getValue(node *ListNode) int {
	if node == nil {
		return 0
	}

	return node.Val
}

func getNext(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}

	return node.Next
}
