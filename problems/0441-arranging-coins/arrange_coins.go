package _441_arranging_coins

func arrangeCoins(n int) int {
	count := 1
	for count <= n {
		n -= count
		count++
	}

	return count - 1
}
