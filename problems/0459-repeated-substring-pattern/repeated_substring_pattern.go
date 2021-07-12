package _459_repeated_substring_pattern

func repeatedSubstringPattern(s string) bool {
	for i := 1; i <= len(s)/2; i++ {
		if check(s, i) {
			return true
		}
	}
	return false
}

func check(s string, i int) bool {
	if len(s)%i != 0 {
		return false
	}

	for j := i; j < len(s); j++ {
		if s[j] != s[j%i] {
			return false
		}
	}
	return true
}
