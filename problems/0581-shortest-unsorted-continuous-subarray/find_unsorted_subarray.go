package _581_shortest_unsorted_continuous_subarray

import "math"

func findUnsortedSubarray(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	var min, max, left, right int
	left, right = -1, -1
	min = math.MaxInt32
	max = math.MinInt32
	for i := 0; i < len(nums); i++ {
		if max <= nums[i] {
			max = nums[i]
		} else {
			right = i
		}
	}
	if right == -1 {
		return 0
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if min >= nums[i] {
			min = nums[i]
		} else {
			left = i
		}
	}
	return right - left + 1
}
