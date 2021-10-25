package _451_sort_characters_by_frequency

import (
	"sort"
	"strings"
)

type entry struct {
	ch    rune
	count int
}

func frequencySort(s string) string {
	entries := [128]entry{}

	for _, ch := range s {
		entries[ch] = entry{ch, entries[ch].count + 1}
	}

	sort.Slice(entries[:], func(i, j int) bool {
		return entries[i].count > entries[j].count
	})

	var sb strings.Builder
	sb.Grow(len(s))
	for i := 0; i < len(entries) && entries[i].count > 0; i++ {
		for j := 0; j < entries[i].count; j++ {
			sb.WriteRune(entries[i].ch)
		}
	}

	return sb.String()
}
