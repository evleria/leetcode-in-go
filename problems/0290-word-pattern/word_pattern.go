package _290_word_pattern

import "strings"

func wordPattern(pattern string, s string) bool {
	strs := strings.Split(s, " ")
	if len(pattern) != len(strs) {
		return false
	}

	mapping, used := [26]string{}, map[string]bool{}
	for i := 0; i < len(pattern); i++ {
		if x := mapping[pattern[i]-'a']; x != "" {
			if strs[i] != x {
				return false
			}
		} else {
			if used[strs[i]] {
				return false
			}

			mapping[pattern[i]-'a'] = strs[i]
			used[strs[i]] = true
		}
	}

	return true
}
