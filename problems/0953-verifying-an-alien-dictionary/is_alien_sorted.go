package _953_verifying_an_alien_dictionary

import (
	"math"
	"sort"
)

func isAlienSorted(words []string, order string) bool {
	mapping := [26]int{}
	for i := 0; i < 26; i++ {
		mapping[order[i]-'a'] = i
	}

	return sort.SliceIsSorted(words, func(i, j int) bool {
		return lessOrEqual(words[i], words[j], &mapping)
	})
}

func lessOrEqual(w1, w2 string, mapping *[26]int) bool {
	l := int(math.Min(float64(len(w1)), float64(len(w2))))
	for j := 0; j < l; j++ {
		comp := mapping[w1[j]-'a'] - mapping[w2[j]-'a']
		if comp != 0 {
			return comp < 0
		}
	}
	return len(w1) < len(w2)
}
