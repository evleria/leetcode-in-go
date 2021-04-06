package _551_minimum_operations_to_make_array_equal

func minOperations(n int) int {
	h := n / 2
	result := h * h
	if n%2 != 0 {
		result += h
	}
	return result
}
