package _451_sort_characters_by_frequency

import (
	"sort"
	"strings"
)

type entry struct {
	ch    byte
	count int
}

func frequencySort(s string) string {
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	entries := make([]entry, 0, len(m))
	for ch, count := range m {
		entries = append(entries, entry{ch, count})
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].count > entries[j].count
	})

	var sb strings.Builder
	sb.Grow(len(s))
	for _, e := range entries {
		for i := 0; i < e.count; i++ {
			sb.WriteByte(e.ch)
		}
	}

	return sb.String()
}
