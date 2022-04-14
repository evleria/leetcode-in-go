package _206_divide_array_into_equal_pairs

import "sort"

func divideArray(nums []int) bool {
	if len(nums)%2 != 0 {
		return false
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		if nums[i] != nums[i+1] {
			return false
		}
	}
	return true
}
