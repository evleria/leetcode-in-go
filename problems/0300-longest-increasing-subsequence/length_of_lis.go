package _300_longest_increasing_subsequence

func lengthOfLIS(nums []int) int {
	slice, result := make([]int, len(nums)+1), 0
	slice[0] = 0
	for i := 1; i <= len(nums); i++ {
		for j := 1; j < i; j++ {
			if nums[j-1] < nums[i-1] {
				slice[i] = max(slice[i], slice[j])
			}
		}
		slice[i] = slice[i] + 1
		result = max(result, slice[i])
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
