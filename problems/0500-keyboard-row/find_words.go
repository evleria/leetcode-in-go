package _500_keyboard_row

var letters = [26]int{2, 3, 3, 2, 1, 2, 2, 2, 1, 2, 2, 2, 3, 3, 1, 1, 1, 1, 2, 1, 1, 3, 1, 3, 1, 3}

func findWords(words []string) []string {
	result := []string{}
	for _, word := range words {
		if inOneRow(word) {
			result = append(result, word)
		}
	}
	return result
}

func inOneRow(word string) bool {
	r := letters[toIndex(word[0])]
	for i := 1; i < len(word); i++ {
		if letters[toIndex(word[i])] != r {
			return false
		}
	}
	return true
}

func toIndex(b byte) byte {
	if 'A' <= b && b <= 'Z' {
		return b - 'A'
	}
	return b - 'a'
}
