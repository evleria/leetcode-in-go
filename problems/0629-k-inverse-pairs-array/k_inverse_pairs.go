package _629_k_inverse_pairs_array

const Mod = 1_000_000_007

func kInversePairs(n int, k int) int {
	if k == 0 {
		return 1
	}

	dp := make([]int, k+1)
	for i := 1; i <= n; i++ {
		temp := make([]int, k+1)
		temp[0] = 1
		for j := 1; j <= k; j++ {
			x := 0
			if j >= i {
				x = dp[j-i]
			}
			temp[j] = (temp[j-1] + dp[j] - x) % Mod
		}
		dp = temp
	}
	return (dp[k] + Mod - dp[k-1]) % Mod
}
