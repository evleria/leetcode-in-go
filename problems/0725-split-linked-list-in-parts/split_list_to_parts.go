package _725_split_linked_list_in_parts

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
)

func splitListToParts(head *ListNode, k int) []*ListNode {
	count := 0
	for cur := head; cur != nil; cur = cur.Next {
		count++
	}

	result := make([]*ListNode, k)
	for i := range result {
		l := count / k
		if count%k > i {
			l++
		}

		result[i] = head
		for j := 0; j < l; j++ {
			if j < l-1 {
				head = head.Next
			} else {
				head, head.Next = head.Next, nil
			}
		}
	}

	return result
}
