package _089_gray_code

func grayCode(n int) []int {
	result := make([]int, 1<<n)
	result[1] = 1

	for i := 1; i < n; i++ {
		offset := 1 << i
		for j := 0; j < offset; j++ {
			result[j+offset] = result[offset-j-1] | offset
		}
	}

	return result
}
