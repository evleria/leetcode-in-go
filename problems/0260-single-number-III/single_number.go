package _260_single_number_III

func singleNumber(nums []int) []int {
	result := make([]int, 2)
	a := 0
	for _, n := range nums {
		a ^= n
	}

	rightBit := a & -a
	b := 0
	for _, n := range nums {
		if n&rightBit != 0 {
			b ^= n
		}
	}
	result[0] = b
	result[1] = a ^ b

	return result
}
