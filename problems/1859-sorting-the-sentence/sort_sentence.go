package _859_sorting_the_sentence

import (
	"sort"
	"strings"
)

func sortSentence(s string) string {
	strs := strings.Split(s, " ")
	sort.Slice(strs, func(i, j int) bool {
		li, lj := len(strs[i]), len(strs[j])
		return strs[i][li-1] < strs[j][lj-1]
	})
	for i := 0; i < len(strs); i++ {
		strs[i] = strs[i][:len(strs[i])-1]
	}
	return strings.Join(strs, " ")
}
