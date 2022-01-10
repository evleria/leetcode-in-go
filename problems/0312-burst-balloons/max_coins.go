package _312_burst_balloons

func maxCoins(nums []int) int {
	val := make([]int, 0, len(nums)+2)
	val = append(val, 1)
	val = append(val, nums...)
	val = append(val, 1)
	n := len(val)

	dp := make([][]int, len(nums)+2)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(dp)+2)
	}

	for l := 3; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := i + l - 1
			for k := i + 1; k < j; k++ {
				sum := dp[i][k] + dp[k][j] + val[i]*val[k]*val[j]
				dp[i][j] = max(dp[i][j], sum)
			}
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
