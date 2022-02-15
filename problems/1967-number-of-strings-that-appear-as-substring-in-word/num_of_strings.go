package _967_number_of_strings_that_appear_as_substring_in_word

import "strings"

func numOfStrings(patterns []string, word string) int {
	count := 0
	for _, pattern := range patterns {
		if strings.Contains(word, pattern) {
			count++
		}
	}
	return count
}
