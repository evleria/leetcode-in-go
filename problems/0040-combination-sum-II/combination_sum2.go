package _040_combination_sum_II

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := make([][]int, 0, 16)
	backtrack(&result, nil, candidates, target)
	return result
}

func backtrack(result *[][]int, buffer []int, candidates []int, target int) {
	if target == 0 {
		cp := make([]int, len(buffer))
		copy(cp, buffer)
		*result = append(*result, cp)
		return
	}
	for i := 0; i < len(candidates); i++ {
		if candidate := candidates[i]; candidate <= target {
			backtrack(result, append(buffer, candidate), candidates[i+1:], target-candidate)
			for ; i < len(candidates)-1 && candidates[i] == candidates[i+1]; i++ {
			}
		}
	}
}
