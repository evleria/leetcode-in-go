package _800_maximum_ascending_subarray_sum

func maxAscendingSum(nums []int) int {
	max, sum := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			sum += nums[i]
		} else {
			max, sum = maximum(max, sum), nums[i]
		}

	}
	max = maximum(max, sum)
	return max
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
