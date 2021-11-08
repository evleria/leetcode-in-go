package _011_final_value_of_variable_after_performing_operations

func finalValueAfterOperations(operations []string) int {
	sum := 0
	for _, operation := range operations {
		if operation == "--X" || operation == "X--" {
			sum--
		}
		if operation == "X++" || operation == "++X" {
			sum++
		}
	}
	return sum
}
