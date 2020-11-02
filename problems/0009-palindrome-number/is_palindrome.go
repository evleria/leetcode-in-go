package _009_palindrome_number

func isPalindrome(x int) bool {
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	}

	y := 0
	for ; x > y; x /= 10 {
		y *= 10
		y += x % 10
	}

	return x == y || x == y/10
}
