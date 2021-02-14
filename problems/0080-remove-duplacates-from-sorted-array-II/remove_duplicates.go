package _080_remove_duplacates_from_sorted_array_II

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	out := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[out-2] {
			nums[out] = nums[i]
			out++
		}
	}
	return out
}
