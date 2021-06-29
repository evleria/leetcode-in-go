package _004_max_consecutive_ones_III

func longestOnes(nums []int, k int) int {
	max, zeroes := 0, make([]int, 0, k)
	for left, right := 0, 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroes = append(zeroes, right)

			if len(zeroes) > k {
				left, zeroes = zeroes[0]+1, zeroes[1:]
			}
		}

		if l := right - left + 1; l > max {
			max = l
		}
	}
	return max
}
