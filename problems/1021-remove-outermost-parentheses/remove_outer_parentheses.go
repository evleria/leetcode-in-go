package _021_remove_outermost_parentheses

import "strings"

func removeOuterParentheses(s string) string {
	var sb strings.Builder
	sb.Grow(len(s))
	depth := 0
	for _, r := range s {
		if r == '(' && depth > 0 || r == ')' && depth > 1 {
			sb.WriteRune(r)
		}
		if r == '(' {
			depth++
		} else {
			depth--
		}
	}
	return sb.String()
}
