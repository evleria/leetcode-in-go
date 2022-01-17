package _114_maximum_number_of_words_found_in_sentences

import "strings"

func mostWordsFound(sentences []string) int {
	result := 0
	for _, sentence := range sentences {
		words := strings.Split(sentence, " ")
		count := len(words)
		if count > result {
			result = count
		}
	}
	return result
}
