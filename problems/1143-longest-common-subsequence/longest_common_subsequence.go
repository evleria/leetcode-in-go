package _143_longest_common_subsequence

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	for i, v1 := range text1 {
		for j, v2 := range text2 {
			dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			if v1 == v2 {
				dp[i+1][j+1] = max(dp[i][j]+1, dp[i+1][j+1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
