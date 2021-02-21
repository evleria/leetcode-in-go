package _213_house_robber_II

func rob(nums []int) int {
	if len(nums) <= 3 {
		return max(nums...)
	}

	var a, b, c, d int
	for i := 0; i < len(nums)-1; i++ {
		a, b, c, d = b, max(b, a+nums[i]), d, max(d, c+nums[i+1])
	}
	return max(b, d)
}

func max(nums ...int) int {
	var result int
	for _, n := range nums {
		if n > result {
			result = n
		}
	}
	return result
}
