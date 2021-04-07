package _704_determine_if_string_halves_are_alike

import (
	"strings"
)

func halvesAreAlike(s string) bool {
	s = strings.ToLower(s)
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	acc := 0
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		if vowels[s[left]] {
			acc++
		}
		if vowels[s[right]] {
			acc--
		}
	}
	return acc == 0
}
