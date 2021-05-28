package _652_defuse_the_bomb

func decrypt(code []int, k int) []int {
	n := len(code)
	result := make([]int, n)
	if k == 0 {
		return result
	}

	l, r := 1, k
	if k < 0 {
		l, r = n+k, n-1
	}

	window := 0
	for i := l; i <= r; i++ {
		window += code[i]
	}

	for offset := 0; offset < len(code); offset++ {
		result[offset] = window
		window += code[(r+offset+1)%n] - code[(l+offset)%n]
	}
	return result
}
