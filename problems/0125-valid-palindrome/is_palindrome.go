package _125_valid_palindrome

import (
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		for ; left < right && !isAlphaNumeric(s[left]); left++ {
		}
		for ; left < right && !isAlphaNumeric(s[right]); right-- {
		}

		if s[left] != s[right] {
			return false
		}
	}
	return true
}

func isAlphaNumeric(a byte) bool {
	return a >= 'a' && a <= 'z' || a >= '0' && a <= '9'
}
