package _047_permutations_II

import "sort"

func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0, 16)
	buffer := make([]int, 0, len(nums))
	sort.Ints(nums)
	backtrack(&result, buffer, nums)
	return result
}

func backtrack(result *[][]int, buffer []int, nums []int) {
	if len(nums) == 0 {
		cp := make([]int, len(buffer))
		copy(cp, buffer)
		*result = append(*result, cp)
		return
	}
	for i := 0; i < len(nums); i++ {
		backtrack(result, append(buffer, nums[i]), append(append([]int{}, nums[:i]...), nums[i+1:]...))
		for ; i < len(nums)-1 && nums[i] == nums[i+1]; i++ {
		}
	}
}
