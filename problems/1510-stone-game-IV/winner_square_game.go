package _510_stone_game_IV

func winnerSquareGame(n int) bool {
	dp := make([]bool, n+1)
	for i := 0; i < n+1; i++ {
		for k := 1; k*k <= i; k++ {
			if !dp[i-k*k] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
