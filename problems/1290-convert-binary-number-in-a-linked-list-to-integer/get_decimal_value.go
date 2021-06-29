package _290_convert_binary_number_in_a_linked_list_to_integer

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
)

func getDecimalValue(head *ListNode) int {
	acc := 0
	for ; head != nil; head = head.Next {
		acc *= 2
		acc += head.Val
	}
	return acc
}
