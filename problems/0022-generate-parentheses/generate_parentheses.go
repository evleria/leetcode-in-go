package _022_generate_parentheses

func generateParenthesis(n int) []string {
	result := make([]string, 0, 16)
	bytes := make([]byte, 0, n*2)
	backtrack(&result, bytes, 0, 0, n)
	return result
}

func backtrack(result *[]string, bytes []byte, open int, closed int, max int) {
	if open == max && closed == max {
		*result = append(*result, string(bytes))
		return
	}

	if open < max {
		backtrack(result, append(bytes, '('), open+1, closed, max)
	}
	if closed < open {
		backtrack(result, append(bytes, ')'), open, closed+1, max)
	}
}
