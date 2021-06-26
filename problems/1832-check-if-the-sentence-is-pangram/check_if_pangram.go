package _832_check_if_the_sentence_is_pangram

func checkIfPangram(sentence string) bool {
	seen, left := [26]bool{}, 26
	for i := 0; i < len(sentence) && left > 0 && len(sentence)-i >= left; i++ {
		if !seen[sentence[i]-'a'] {
			seen[sentence[i]-'a'] = true
			left--
		}
	}
	return left == 0
}
