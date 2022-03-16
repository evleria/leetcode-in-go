package _946_validate_stack_sequences

func validateStackSequences(pushed []int, popped []int) bool {
	result := []int{}
	for i := 0; i < len(pushed); i++ {
		result = append(result, pushed[i])
		for len(result) > 0 && result[len(result)-1] == popped[0] {
			result, popped = result[:len(result)-1], popped[1:]
		}
	}
	return len(result) <= 0
}
