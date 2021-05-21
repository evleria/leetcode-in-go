package _662_check_if_two_string_arrays_are_equivalent

import "strings"

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	if strings.Join(word1, "") == strings.Join(word2, "") {
		return true
	}
	return false
}
