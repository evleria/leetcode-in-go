package _062_unique_paths

func uniquePaths(m int, n int) int {
	row := make([]int, n)
	row[0] = 1

	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			row[j] = row[j-1] + row[j]
		}
	}

	return row[n-1]
}
