package _583_delete_operation_for_two_strings

func minDistance(word1 string, word2 string) int {
	dp := make([]int, len(word2))
	for i := 0; i < len(word1); i++ {
		var prev int
		for j := 0; j < len(word2); j++ {
			temp := dp[j]
			if word1[i] == word2[j] {
				dp[j] = prev + 1
			} else if j > 0 && dp[j-1] > dp[j] {
				dp[j] = dp[j-1]
			}
			prev = temp
		}
	}
	return len(word1) + len(word2) - 2*dp[len(word2)-1]
}
