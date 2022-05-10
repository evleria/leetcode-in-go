package _216_combination_sum_III

func combinationSum3(k int, n int) [][]int {
	result := [][]int{}
	buffer := make([]int, 0, k)

	backtrack(buffer, n, &result)
	return result
}

func backtrack(buffer []int, left int, result *[][]int) {
	if cap(buffer) == len(buffer) {
		if left == 0 {
			cp := make([]int, len(buffer))
			copy(cp, buffer)
			*result = append(*result, cp)
		}
		return
	}

	last := 0
	if len(buffer) > 0 {
		last = buffer[len(buffer)-1]
	}

	for i := last + 1; i <= 9 && i <= left; i++ {
		backtrack(append(buffer, i), left-i, result)
	}
}
