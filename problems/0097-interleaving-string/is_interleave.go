package _097_interleaving_string

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make([]bool, len(s1)+1)
	for i := 0; i < len(s2)+1; i++ {
		for j := 0; j < len(s1)+1; j++ {
			dp[j] = i == 0 && j == 0 || j > 0 && dp[j-1] && s1[j-1] == s3[i+j-1] || i > 0 && dp[j] && s2[i-1] == s3[i+j-1]
		}
	}

	return dp[len(s1)]
}
