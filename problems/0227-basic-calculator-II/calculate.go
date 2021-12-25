package _227_basic_calculator_II

func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}
	currNum, res := 0, 0
	stack := make([]int, 0)
	var ops byte = '+'

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && i < len(s)-1 {
			continue
		}

		if s[i] >= '0' && s[i] <= '9' {
			currNum = currNum*10 + int(s[i]-'0')
			if i < len(s)-1 {
				continue
			}
		}

		if ops == '+' {
			stack = append(stack, currNum)
		}

		if ops == '-' {
			stack = append(stack, -currNum)
		}

		if ops == '*' {
			temp := stack[len(stack)-1] * currNum
			stack = stack[:len(stack)-1]
			stack = append(stack, temp)
		}

		if ops == '/' {
			temp := stack[len(stack)-1] / currNum
			stack = stack[:len(stack)-1]
			stack = append(stack, temp)
		}
		currNum = 0
		ops = s[i]
	}

	for _, v := range stack {
		res += v
	}

	return res
}
