package _304_find_n_unique_integers_sum_up_to_zero

func sumZero(n int) []int {
	result := make([]int, n)
	for i := 1; i < n; i++ {
		result[i] = i
	}
	result[0] = -n * (n - 1) / 2
	return result
}
