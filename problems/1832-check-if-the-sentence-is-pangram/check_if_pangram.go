package _832_check_if_the_sentence_is_pangram

func checkIfPangram(sentence string) bool {
	if len(sentence) < 26 {
		return false
	}
	seen, count := [26]bool{}, 0
	for i := 0; i < len(sentence) && count < 26; i++ {
		if !seen[sentence[i]-'a'] {
			seen[sentence[i]-'a'] = true
			count++
		}
	}
	return count == 26
}
