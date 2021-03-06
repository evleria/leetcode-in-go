package _820_short_encoding_of_words

func minimumLengthEncoding(words []string) int {
	m := make(map[string]bool, len(words))
	for _, word := range words {
		m[word] = true
	}
	for word := range m {
		for i := 1; i < len(word); i++ {
			delete(m, word[i:])
		}
	}
	result := 0
	for word := range m {
		result += len(word) + 1
	}
	return result
}
