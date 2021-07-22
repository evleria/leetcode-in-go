package _915_partition_array_into_disjoint_intervals

func partitionDisjoint(nums []int) int {
	leftMax, peak, result := nums[0], nums[0], 1
	for i := 0; i < len(nums); i++ {
		if nums[i] < leftMax {
			result = i + 1
			leftMax = peak
		} else {
			peak = max(peak, nums[i])
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
