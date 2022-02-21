package _169_count_operations_to_obtain_zero

func countOperations(num1 int, num2 int) int {
	count := 0
	for num1 != 0 && num2 != 0 {
		switch {
		case num1 < num2:
			num2 -= num1
		case num1 >= num2:
			num1 -= num2
		}
		count++
	}
	return count
}
