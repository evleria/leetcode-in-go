package _421_maximum_xor_of_two_numbers_in_an_array

func findMaximumXOR(nums []int) int {
	xor := 0
	// try from MSB->LSB
	for i := 31; i >= 0; i-- {
		hash := map[int]bool{}
		for _, v := range nums {
			// set lower i-bit 0
			hash[v>>i<<i] = true
		}
		tryXOR := xor + 1<<i
		for k := range hash {
			if hash[k^tryXOR] {
				// exist, i-bit is valid and do next bit
				xor = tryXOR
				break
			}
		}
	}
	return xor
}
