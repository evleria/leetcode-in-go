package _890_find_and_replace_pattern

func findAndReplacePattern(words []string, pattern string) []string {
	result := []string{}
	for _, word := range words {
		if areIsometric(word, pattern) {
			result = append(result, word)
		}
	}
	return result
}

func areIsometric(s1, s2 string) bool {
	m, used := make(map[byte]byte, len(s1)), make(map[byte]bool, len(s1))
	for i := 0; i < len(s1); i++ {
		if j, ok := m[s1[i]]; ok {
			if s2[i] != j {
				return false
			}
		} else {
			if used[s2[i]] {
				return false
			}
			m[s1[i]], used[s2[i]] = s2[i], true
		}
	}
	return true
}
