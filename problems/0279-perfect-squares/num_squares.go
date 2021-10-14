package _279_perfect_squares

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt64
	}
	dp[0] = 0
	count := 1

	for count*count <= n {
		sq := count * count
		for i := sq; i < n+1; i++ {
			dp[i] = min(dp[i-sq]+1, dp[i])
		}
		count++
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
