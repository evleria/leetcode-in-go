package _020_valid_parentheses

func isValid(s string) bool {
	if s == "" {
		return true
	}
	dict := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := make([]rune, 0, len(s)/2)

	for _, ch := range s {
		if x, ok := dict[ch]; ok {
			stack = append(stack, x)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != ch {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
