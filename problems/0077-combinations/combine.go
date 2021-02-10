package _077_combinations

func combine(n int, k int) [][]int {
	var result [][]int
	buffer := make([]int, 0, k)
	backtrack(&result, buffer, k, n, 1)
	return result
}

func backtrack(result *[][]int, buffer []int, left, n, s int) {
	if left == 0 {
		cp := make([]int, len(buffer))
		copy(cp, buffer)
		*result = append(*result, cp)
		return
	}

	for i := s; i <= n-left+1; i++ {
		backtrack(result, append(buffer, i), left-1, n, i+1)
	}
}
