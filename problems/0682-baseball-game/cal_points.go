package _682_baseball_game

import "strconv"

func calPoints(ops []string) int {
	stack := make([]int, 0, len(ops))

	for i := 0; i < len(ops); i++ {
		switch ops[i] {
		case "+":
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
		case "D":
			stack = append(stack, stack[len(stack)-1]*2)
		case "C":
			stack = stack[:len(stack)-1]
		default:
			v, _ := strconv.Atoi(ops[i])
			stack = append(stack, v)
		}
	}

	sum := 0
	for _, n := range stack {
		sum += n
	}
	return sum
}
