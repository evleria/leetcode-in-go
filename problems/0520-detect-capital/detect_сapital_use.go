package _520_detect_capital

func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}

	var min, max byte = 'a', 'z'
	if isUpper(word[0]) && isUpper(word[1]) {
		min, max = 'A', 'Z'
	}

	for i := 1; i < len(word); i++ {
		if word[i] < min || word[i] > max {
			return false
		}
	}

	return true
}

func isUpper(ch byte) bool {
	return ch >= 'A' && ch <= 'Z'
}
