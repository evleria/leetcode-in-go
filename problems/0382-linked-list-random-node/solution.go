package _382_linked_list_random_node

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	"math/rand"
)

type Solution struct {
	nums []int
}

func Constructor(head *ListNode) Solution {
	var nums []int
	for ; head != nil; head = head.Next {
		nums = append(nums, head.Val)
	}

	return Solution{nums: nums}
}

func (s *Solution) GetRandom() int {
	return s.nums[rand.Intn(len(s.nums))]
}
