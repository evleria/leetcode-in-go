package _600_non_negative_integers_without_consecutive_ones

func findIntegers(n int) int {
	x, y, result := 1, 2, 0
	n++
	for n != 0 {
		if n&0b11 == 0b11 {
			result = 0
		}
		result += x * (n & 1)
		n >>= 1
		x, y = y, x+y
	}

	return result
}
