package _322_coin_change

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		min := amount + 1
		for _, coin := range coins {
			if x := i - coin; x >= 0 {
				if r := dp[x] + 1; r < min {
					min = r
				}
			}
		}
		dp[i] = min
	}
	if result := dp[amount]; result != amount+1 {
		return result
	}
	return -1
}
