package _318_maximum_product_of_word_lengths

func maxProduct(words []string) int {
	max, chars := 0, make([]uint32, len(words))
	for i, word := range words {
		for j := 0; j < len(word); j++ {
			chars[i] |= 1 << (word[j] - 'a')
		}
	}

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if product := len(words[i]) * len(words[j]); chars[i]&chars[j] == 0 && product > max {
				max = product
			}
		}
	}

	return max
}
