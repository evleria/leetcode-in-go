package _532_k_diff_pairs_in_an_array

import "sort"

func findPairs(nums []int, k int) int {
	count := 0
	sort.Ints(nums)
	j := 0
	for i := 0; i < len(nums)-1; i++ {
		for ; j <= i || j < len(nums) && nums[j] < nums[i]+k; j++ {
		}
		if j == len(nums) {
			break
		} else if nums[j] == nums[i]+k {
			count++
			for ; i < len(nums)-1 && nums[i] == nums[i+1]; i++ {
			}
		}
	}
	return count
}
