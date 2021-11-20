package _540_single_element_in_a_sorted_array

func singleNonDuplicate(nums []int) int {
	res := 0
	for i := range nums {
		res ^= nums[i]
	}
	return res
}
