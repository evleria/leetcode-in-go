package _260_shift_2d_grid

func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	k %= m * n
	if k == 0 {
		return grid
	}

	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			idx := (i*n + j + k) % (m * n)
			ni, nj := idx/n, idx%n
			res[ni][nj] = grid[i][j]
		}
	}

	return res
}
