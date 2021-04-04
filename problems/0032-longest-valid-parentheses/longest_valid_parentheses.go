package _032_longest_valid_parentheses

func longestValidParentheses(s string) int {
	res, start, n := 0, 0, len(s)
	stack := []int{}
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			if len(stack) == 0 {
				start = i + 1
			} else {
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					res = max(res, i-stack[len(stack)-1])
				} else {
					res = max(res, i-start+1)
				}
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
