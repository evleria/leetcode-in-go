package _139_word_break

import "strings"

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}

	starts := make([]bool, len(s))
	starts[0] = true

	for i := range s {
		if starts[i] {
			for _, prefix := range wordDict {
				if strings.HasPrefix(s[i:], prefix) {
					if nextPoint := i + len(prefix); nextPoint == len(s) {
						return true
					} else {
						starts[nextPoint] = true
					}
				}
			}
		}
	}

	return false
}
