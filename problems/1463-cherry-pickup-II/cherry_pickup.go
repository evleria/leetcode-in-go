package _463_cherry_pickup_II

func cherryPickup(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	dp := make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, n)
		}
	}

	for row := m - 1; row >= 0; row-- {
		for col1 := 0; col1 < n; col1++ {
			for col2 := 0; col2 < n; col2++ {
				result := 0
				// current cell
				result += grid[row][col1]
				if col1 != col2 {
					result += grid[row][col2]
				}
				// transition
				if row != m-1 {
					max := 0
					for newCol1 := col1 - 1; newCol1 <= col1+1; newCol1++ {
						for newCol2 := col2 - 1; newCol2 <= col2+1; newCol2++ {
							if newCol1 >= 0 && newCol1 < n && newCol2 >= 0 && newCol2 < n {
								if dp[row+1][newCol1][newCol2] > max {
									max = dp[row+1][newCol1][newCol2]
								}
							}
						}
					}
					result += max
				}
				dp[row][col1][col2] = result
			}
		}
	}
	return dp[0][0][n-1]
}
