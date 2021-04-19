package _377_combination_sum_IV

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for t := 1; t <= target; t++ {
		for i := 0; i < len(nums); i++ {
			if nums[i] <= t {
				dp[t] += dp[t-nums[i]]
			}
		}
	}
	return dp[target]
}
