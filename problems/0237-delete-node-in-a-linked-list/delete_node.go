package _237_delete_node_in_a_linked_list

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
)

func deleteNode(node *ListNode) {
	for node.Next != nil {
		node.Val = node.Next.Val

		if node.Next.Next != nil {
			node = node.Next
		} else {
			node.Next = nil
		}
	}
}
