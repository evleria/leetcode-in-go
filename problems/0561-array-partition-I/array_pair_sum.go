package _561_array_partition_I

import "sort"

func arrayPairSum(nums []int) int {
	ans := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		ans += min(nums[i], nums[i+1])
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
