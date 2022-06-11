package _658_minimum_operations_to_reduce_x_to_zero

func minOperations(nums []int, x int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	maxi := -1
	l, cur := 0, 0

	for r := 0; r < len(nums); r++ {
		cur += nums[r]
		for ; cur > total-x && l <= r; l++ {
			cur -= nums[l]
		}

		if cur == total-x {
			maxi = max(maxi, r-l+1)
		}
	}
	if maxi != -1 {
		return len(nums) - maxi
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
