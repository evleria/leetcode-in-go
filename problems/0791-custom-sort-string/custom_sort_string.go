package _791_custom_sort_string

import "strings"

func customSortString(order string, str string) string {
	m := [26]int{}
	for _, r := range str {
		m[r-'a']++
	}

	var sb strings.Builder
	sb.Grow(len(str))
	for _, r := range order {
		sb.WriteString(strings.Repeat(string(r), m[r-'a']))
		m[r-'a'] = 0
	}
	for i, c := range m {
		sb.WriteString(strings.Repeat(string(rune('a'+i)), c))
	}
	return sb.String()
}
