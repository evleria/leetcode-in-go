package _764_largest_plus_sign

func orderOfLargestPlusSign(n int, mines [][]int) int {
	zeros := make(map[[2]int]bool, len(mines))
	for _, mine := range mines {
		zeros[[2]int{mine[0], mine[1]}] = true
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	helper := func(count *int, i, j int) int {
		if zeros[[2]int{i, j}] {
			*count = 0
		} else {
			*count++
		}

		return *count
	}

	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			dp[i][j] = helper(&count, i, j)
		}

		count = 0
		for j := n - 1; j >= 0; j-- {
			dp[i][j] = min(helper(&count, i, j), dp[i][j])
		}
	}

	max := 0
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			dp[j][i] = min(helper(&count, j, i), dp[j][i])
		}

		count = 0
		for j := n - 1; j >= 0; j-- {
			x := min(helper(&count, j, i), dp[j][i])
			if x > max {
				max = x
			}
		}
	}

	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
