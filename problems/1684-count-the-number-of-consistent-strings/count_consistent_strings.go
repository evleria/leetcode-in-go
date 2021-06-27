package _684_count_the_number_of_consistent_strings

func countConsistentStrings(allowed string, words []string) int {
	letters := [26]bool{}
	for _, ch := range allowed {
		letters[ch-'a'] = true
	}
	isAllowed := func(word string) bool {
		for i := 0; i < len(word); i++ {
			if !letters[word[i]-'a'] {
				return false
			}
		}
		return true
	}

	count := 0
	for _, ch := range words {
		if isAllowed(ch) {
			count++
		}
	}
	return count
}
