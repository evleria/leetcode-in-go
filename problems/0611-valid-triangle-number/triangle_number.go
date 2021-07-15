package _611_valid_triangle_number

import "sort"

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	count := 0
	for k := len(nums) - 1; k > 1; k-- {
		for i, j := 0, k-1; i < j; {
			if nums[i]+nums[j] > nums[k] {
				count += j - i
				j--
			} else {
				i++
			}
		}
	}
	return count
}
