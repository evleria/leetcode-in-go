package _462_minimum_moves_to_equal_array_elements_II

import "sort"

func minMoves2(nums []int) int {
	sort.Ints(nums)
	count := 0
	for left, right := 0, len(nums)-1; left < right; left, right = left+1, right-1 {
		count += nums[right] - nums[left]
	}
	return count
}
