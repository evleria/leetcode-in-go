package _557_reverse_words_in_a_string_III

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i, word := range words {
		words[i] = reverse(word)
	}
	return strings.Join(words, " ")
}

func reverse(word string) string {
	str := []byte(word)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	return string(str)
}
