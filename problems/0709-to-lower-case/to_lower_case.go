package _709_to_lower_case

import "strings"

func toLowerCase(str string) string {
	hasUpper := false
	for i := 0; i < len(str) && !hasUpper; i++ {
		hasUpper = hasUpper || (str[i] >= 'A' && str[i] <= 'Z')
	}

	if !hasUpper {
		return str
	}

	var b strings.Builder
	b.Grow(len(str))
	for i := 0; i < len(str); i++ {
		ch := str[i]
		if ch >= 'A' && ch <= 'Z' {
			ch += 'a' - 'A'
		}
		b.WriteByte(ch)
	}
	return b.String()
}
