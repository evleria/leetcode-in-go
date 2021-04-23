package _696_count_binary_substrings

import "math"

func countBinarySubstrings(s string) int {
	count, prevCount, result := 1, 0, 0
	for i := 1; i <= len(s); i++ {
		if i < len(s) && s[i] == s[i-1] {
			count++
		} else {
			result += int(math.Min(float64(count), float64(prevCount)))
			count, prevCount = 1, count
		}
	}

	return result
}
