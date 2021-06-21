package _342_number_of_steps_to_reduce_a_number_to_zero

func numberOfSteps(num int) int {
	count := 0
	for num != 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num--
		}
		count++
	}
	return count
}
