package _147_insertion_sort_list

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
)

func insertionSortList(head *ListNode) *ListNode {
	dummy := new(ListNode)
	for head != nil {
		cur := dummy
		for ; cur.Next != nil && cur.Next.Val < head.Val; cur = cur.Next {
		}
		cur.Next, head.Next, head = head, cur.Next, head.Next
	}
	return dummy.Next
}
