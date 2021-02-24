package _241_different_ways_to_add_parentheses

import "strconv"

var operations = map[byte]func(int, int) int{
	'+': func(a, b int) int { return a + b },
	'-': func(a, b int) int { return a - b },
	'*': func(a, b int) int { return a * b },
}

func diffWaysToCompute(input string) []int {
	var result []int
	for i := 0; i < len(input); i++ {
		if op, ok := operations[input[i]]; ok {
			left, right := diffWaysToCompute(input[:i]), diffWaysToCompute(input[i+1:])
			for _, l := range left {
				for _, r := range right {
					result = append(result, op(l, r))
				}
			}
		}
	}
	if len(result) == 0 {
		n, _ := strconv.Atoi(input)
		result = append(result, n)
	}

	return result
}
