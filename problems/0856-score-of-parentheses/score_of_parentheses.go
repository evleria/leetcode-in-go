package _856_score_of_parentheses

func scoreOfParentheses(s string) int {
	str := []byte(s)
	result, open := 0, 0
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			open++
		} else {
			open--
			if str[i-1] == '(' {
				result += 1 << open
			}
		}
	}
	return result
}
