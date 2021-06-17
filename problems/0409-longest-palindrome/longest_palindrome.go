package _409_longest_palindrome

func longestPalindrome(s string) int {
	hasPair := [58]bool{}
	result := 0
	for i := 0; i < len(s); i++ {
		idx := s[i] - 'A'
		if hasPair[idx] {
			result += 2
			hasPair[idx] = false
		} else {
			hasPair[idx] = true
		}
	}
	if result < len(s) {
		result++
	}
	return result
}
