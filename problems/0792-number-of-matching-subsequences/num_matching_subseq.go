package _792_number_of_matching_subsequences

func numMatchingSubseq(s string, words []string) int {
	awaiting := [26][]string{}
	insert := func(s string) {
		idx := s[0] - 'a'
		awaiting[idx] = append(awaiting[idx], s[1:])
	}

	for _, word := range words {
		insert(word)
	}

	count := 0
	for i := 0; i < len(s) && count < len(words); i++ {
		idx := s[i] - 'a'
		for _, w := range awaiting[idx] {
			awaiting[idx] = awaiting[idx][1:]
			if len(w) == 0 {
				count++
			} else {
				insert(w)
			}
		}
	}

	return count
}
