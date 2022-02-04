package _525_contiguous_array

func findMaxLength(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = dp[i-1] + nums[i-1]*2 - 1
	}
	for d := n >> 1 << 1; d > 0; d -= 2 {
		for s := 0; s+d <= n; s++ {
			if dp[s] == dp[s+d] {
				return d
			}
		}
	}
	return 0
}
