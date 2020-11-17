package _005_longest_palindromic_substring

func longestPalindrome(s string) string {
	result := ""

	for start := 0; start < len(s); start++ {
		left, right := start, start

		for right < len(s)-1 && s[right] == s[right+1] {
			right++
			start++
		}

		for left > 0 && right < len(s)-1 && s[left-1] == s[right+1] {
			left--
			right++
		}

		if right-left+1 > len(result) {
			result = s[left : right+1]
		}
	}

	return result
}
