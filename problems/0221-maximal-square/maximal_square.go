package _221_maximal_square

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	side := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			dp[i+1][j+1] = min(min(dp[i][j+1], dp[i+1][j]), dp[i][j]) + 1
			if dp[i+1][j+1] > side {
				side = dp[i+1][j+1]
			}
		}
	}
	return side * side
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
