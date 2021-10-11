package _728_self_dividing_numbers

func selfDividingNumbers(left int, right int) []int {
	result := []int{}
	for i := left; i <= right; i++ {
		if isSelfDividing(i) {
			result = append(result, i)
		}
	}

	return result
}

func isSelfDividing(n int) bool {
	for i := n; i > 0; i /= 10 {
		d := i % 10
		if d == 0 || n%d != 0 {
			return false
		}
	}
	return true
}
