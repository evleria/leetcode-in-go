package _485_max_consecutive_ones

func findMaxConsecutiveOnes(nums []int) int {
	count, max := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
			if count > max {
				max = count
			}
		} else {
			count = 0
		}
	}

	return max
}
