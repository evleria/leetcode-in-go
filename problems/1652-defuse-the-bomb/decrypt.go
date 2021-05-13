package _652_defuse_the_bomb

// 12345 | k = 3 | 9(234)  12(345) 10(451) 8(512) 6(123)
// 345   | k = -3| 12(345) 9(234)  6(123)  8(512) 10(451)
// -5 +2

func decrypt(code []int, k int) []int {
	result := make([]int, len(code))
	if k == 0 {
		return result
	}
	if k > 0 {
		window := 0
		for i := 1; i <= k; i++ {
			window += code[i]
		}
		result[0] = window

		for i := 1; i < len(code); i++ {
			window += code[(k+i)%len(code)] - code[i]
			result[i] = window
		}
	}
	return result
}
