package _178_number_of_valid_words_for_each_puzzle

func findNumOfValidWords(words []string, puzzles []string) []int {
	freqWords := make(map[int]int)
	for _, word := range words {
		bs := bits(word)
		if _, ok := freqWords[bs]; !ok {
			freqWords[bs] = 1
		} else {
			freqWords[bs]++
		}
	}

	result := make([]int, len(puzzles))
	for i, puzzle := range puzzles {
		mask, num, fb := bits(puzzle), 0, 1<<(puzzle[0]-'a')
		for cur := mask; cur != 0; cur = ((cur - 1) & mask) {
			if (cur & fb) == 0 {
				continue
			}
			if v, ok := freqWords[cur]; ok {
				num += v
			}
		}
		result[i] = num
	}
	return result
}

func bits(word string) int {
	var bits int
	for _, l := range word {
		bits |= 1 << (l - 'a')
	}
	return bits
}
