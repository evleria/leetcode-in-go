package _394_decode_string

import (
	"strings"
	"unicode"
)

func decodeString(s string) string {
	var sb strings.Builder

	var times int
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		switch {
		case unicode.IsLetter(r):
			sb.WriteRune(r)
		case unicode.IsDigit(r):
			times = times*10 + int(r-'0')
		case r == '[':
			end := findEnd(s, i)
			inner := decodeString(s[i+1 : end])
			for t := 0; t < times; t++ {
				sb.WriteString(inner)
			}
			i, times = end, 0
		}
	}

	return sb.String()
}

func findEnd(s string, start int) int {
	level := 1
	for i := start + 1; i < len(s); i++ {
		if s[i] == '[' {
			level++
		} else if s[i] == ']' {
			level--
			if level == 0 {
				return i
			}
		}
	}
	return 0
}
