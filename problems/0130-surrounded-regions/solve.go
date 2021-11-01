package _130_surrounded_regions

var directions = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func solve(board [][]byte) {
	m, n := len(board), len(board[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		board[i][j] = '#'
		for _, dir := range directions {
			x, y := i+dir[0], j+dir[1]
			if x >= 0 && x < m && y >= 0 && y < n && board[x][y] == 'O' {
				dfs(x, y)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i > 0 && i < m-1) && (j > 0 && j < n-1) {
				continue
			}
			if board[i][j] == 'O' {
				dfs(i, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch board[i][j] {
			case 'O':
				board[i][j] = 'X'
			case '#':
				board[i][j] = 'O'
			}
		}
	}
}
