package _338_counting_bits

func countBits(n int) []int {
	result := make([]int, n+1)
	offset := 1
	for i := 1; i <= n; i++ {
		if i == 2*offset {
			offset = i
		}
		result[i] = result[i-offset] + 1
	}
	return result
}
