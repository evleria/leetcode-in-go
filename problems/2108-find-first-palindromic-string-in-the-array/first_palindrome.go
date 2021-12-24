package _108_find_first_palindromic_string_in_the_array

func firstPalindrome(words []string) string {
	for _, word := range words {
		if isPalindrome(word) {
			return word
		}
	}
	return ""
}

func isPalindrome(word string) bool {
	for i := 0; i < len(word); i++ {
		if word[i] != word[len(word)-i-1] {
			return false
		}
	}
	return true
}
