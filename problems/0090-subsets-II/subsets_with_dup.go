package _090_subsets_II

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	buffer := make([]int, 0, len(nums))
	backtrack(buffer, nums, &result)
	return result
}

func backtrack(buffer, rest []int, result *[][]int) {
	cp := make([]int, len(buffer))
	copy(cp, buffer)
	*result = append(*result, cp)

	for i := 0; i < len(rest); i++ {
		backtrack(append(buffer, rest[i]), rest[i+1:], result)
		for ; i < len(rest)-1 && rest[i] == rest[i+1]; i++ {
		}
	}
}
