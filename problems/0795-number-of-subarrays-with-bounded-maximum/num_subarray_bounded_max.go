package _795_number_of_subarrays_with_bounded_maximum

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	j, count, result := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= left && nums[i] <= right {
			result += i - j + 1
			count = i - j + 1
		} else if nums[i] < left {
			result += count
		} else {
			j = i + 1
			count = 0
		}
	}
	return result
}
