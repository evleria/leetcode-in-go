package _046_permutations

func permute(nums []int) [][]int {
	result := make([][]int, 0, 16)
	buffer := make([]int, 0, len(nums))
	backtrack(&result, buffer, nums)

	return result
}

func backtrack(result *[][]int, buffer []int, candidates []int) {
	if len(candidates) == 0 {
		cp := make([]int, len(buffer))
		copy(cp, buffer)
		*result = append(*result, cp)
	} else {
		for i, candidate := range candidates {
			backtrack(result, append(buffer, candidate), append(append([]int{}, candidates[:i]...), candidates[i+1:]...))
		}
	}
}
