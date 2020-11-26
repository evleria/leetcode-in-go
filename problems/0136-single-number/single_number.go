package _136_single_number

func singleNumber(nums []int) int {
	result := 0
	for _, n := range nums {
		result ^= n
	}

	return result
}
