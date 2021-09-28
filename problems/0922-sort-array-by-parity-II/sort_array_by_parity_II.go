package _922_sort_array_by_parity_II

func sortArrayByParityII(nums []int) []int {
	even, odd := 0, 1
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			result[even] = nums[i]
			even += 2
		} else {
			result[odd] = nums[i]
			odd += 2
		}
	}

	return result
}
