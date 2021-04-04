package _474_ones_and_zeroes

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for _, s := range strs {
		z, o := convert(s)
		for i := m; i >= z; i-- {
			for j := n; j >= o; j-- {
				if y := dp[i-z][j-o] + 1; y > dp[i][j] {
					dp[i][j] = y
				}
			}
		}
	}
	return dp[m][n]
}

func convert(s string) (int, int) {
	x := [2]int{}
	for i := 0; i < len(s); i++ {
		x[s[i]-'0']++
	}
	return x[0], x[1]
}
