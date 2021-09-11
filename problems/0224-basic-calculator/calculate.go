package _224_basic_calculator

func calculate(s string) int {
	result, num, op := 0, 0, add
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ' ':
		case '+':
			result, num = op(result, num), 0
			op = add
		case '-':
			result, num = op(result, num), 0
			op = sub
		case '(':
			end := findMatch(s, i)
			num, i = calculate(s[i+1:end]), end
		default:
			num *= 10
			num += int(s[i] - '0')
		}
	}

	result, num = op(result, num), 0
	return result
}

func findMatch(s string, start int) int {
	level := 0
	for i := start; i < len(s); i++ {
		switch s[i] {
		case '(':
			level++
		case ')':
			level--
			if level == 0 {
				return i
			}
		}
	}
	return len(s) - 1
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}
