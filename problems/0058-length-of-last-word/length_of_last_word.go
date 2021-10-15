package _058_length_of_last_word

func lengthOfLastWord(s string) int {
	i, count := len(s)-1, 0
	for ; i >= 0 && s[i] == ' '; i-- {
	}
	for ; i >= 0 && s[i] != ' '; i-- {
		count++
	}
	return count
}
