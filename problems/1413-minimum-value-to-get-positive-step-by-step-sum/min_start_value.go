package _413_minimum_value_to_get_positive_step_by_step_sum

func minStartValue(nums []int) int {
	sum, min := 0, 0
	for _, n := range nums {
		sum += n
		if sum < min {
			min = sum
		}
	}
	return 1 - min
}
