package _201_bitwise_and_of_numbers_range

func rangeBitwiseAnd(left int, right int) int {
	result, mask := 0, 1
	for right > 0 {
		if left == right && left%2 == 1 {
			result |= mask
		}
		left, right, mask = left>>1, right>>1, mask<<1
	}

	return result
}
