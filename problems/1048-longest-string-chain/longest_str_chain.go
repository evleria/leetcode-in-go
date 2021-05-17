package _048_longest_string_chain

import (
	"sort"
)

func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	max, m := 0, make(map[string]int, len(words))
	for _, word := range words {
		l := 1
		for i := 0; i < len(word); i++ {
			if k := m[word[:i]+word[i+1:]] + 1; k > l {
				l = k
			}
		}
		if l > max {
			max = l
		}
		m[word] = l
	}
	return max
}
