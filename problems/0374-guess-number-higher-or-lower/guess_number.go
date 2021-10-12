package _374_guess_number_higher_or_lower

var Pick int

func guess(n int) int {
	switch {
	case n > Pick:
		return -1
	case n < Pick:
		return 1
	default:
		return 0
	}
}

func guessNumber(n int) int {
	start, end := 1, n
	for start <= end {
		mid := (start + end) / 2
		switch guess(mid) {
		case -1:
			end = mid - 1
		case 1:
			start = mid + 1
		default:
			return mid
		}
	}
	return 0
}
