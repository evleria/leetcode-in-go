package _486_xor_operation_in_an_array

func xorOperation(n int, start int) int {
	result := 0
	for i := 0; i < n; i++ {
		result ^= start + 2*i
	}
	return result
}
