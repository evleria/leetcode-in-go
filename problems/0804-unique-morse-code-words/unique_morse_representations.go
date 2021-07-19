package _804_unique_morse_code_words

import "strings"

var letters = [26]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	transformations := map[string]bool{}
	for _, word := range words {
		var sb strings.Builder
		for _, r := range word {
			sb.WriteString(letters[r-'a'])
		}
		transformations[sb.String()] = true
	}
	return len(transformations)
}
