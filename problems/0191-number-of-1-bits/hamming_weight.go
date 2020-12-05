package _191_number_of_1_bits

func hammingWeight(num uint32) int {
	result := 0
	for ; num != 0; num >>= 1 {
		result += int(num & 1)
	}
	return result
}
