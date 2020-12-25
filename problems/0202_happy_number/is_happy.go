package _202_happy_number

func isHappy(n int) bool {
	hare, tortoise := n, n
	for {
		hare, tortoise = squareSum(squareSum(hare)), squareSum(tortoise)
		if hare == 1 || tortoise == 1 {
			return true
		}
		if hare == tortoise {
			return false
		}
	}
}

func squareSum(n int) int {
	acc := 0
	for ; n > 0; n /= 10 {
		d := n % 10
		acc += d * d
	}
	return acc
}
