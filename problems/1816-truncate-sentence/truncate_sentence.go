package _816_truncate_sentence

func truncateSentence(s string, k int) string {
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			k--
			if k == 0 {
				return s[:i]
			}
		}
	}

	return s
}
