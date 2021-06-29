package _67_valid_perfect_square

func isPerfectSquare(num int) bool {
	start, end := 1, num
	for start <= end {
		mid := (start + end) / 2
		sqr := mid * mid
		switch {
		case sqr < num:
			start = mid + 1
		case sqr > num:
			end = mid - 1
		default:
			return true
		}
	}
	return false
}
