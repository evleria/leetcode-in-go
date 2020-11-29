package _053_maximum_subarray

func maxSubArray(nums []int) int {
	max, cur := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if cur+nums[i] > nums[i] {
			cur += nums[i]
		} else {
			cur = nums[i]
		}

		if cur > max {
			max = cur
		}
	}

	return max
}
