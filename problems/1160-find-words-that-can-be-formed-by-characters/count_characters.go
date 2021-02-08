package _160_find_words_that_can_be_formed_by_characters

func countCharacters(words []string, chars string) int {
	letters := [26]int{}
	for i := 0; i < len(chars); i++ {
		letters[chars[i]-'a']++
	}

	result := 0
	for _, word := range words {
		result += countWord(letters, word)
	}
	return result
}

func countWord(letters [26]int, word string) int {
	for i := 0; i < len(word); i++ {
		if letters[word[i]-'a'] == 0 {
			return 0
		}

		letters[word[i]-'a']--
	}

	return len(word)
}
