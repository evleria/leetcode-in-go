package _039_combination_sum

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0, 16)
	current := make([]int, 0, 16)
	backtrack(&result, current, candidates, target)
	return result
}

func backtrack(result *[][]int, current []int, candidates []int, target int) {
	if target == 0 {
		r := make([]int, len(current))
		copy(r, current)
		*result = append(*result, r)
		return
	}

	for i, candidate := range candidates {
		if candidate <= target {
			backtrack(result, append(current, candidate), candidates[i:], target-candidate)
		}
	}
}
