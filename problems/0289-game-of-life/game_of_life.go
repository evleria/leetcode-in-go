package _289_game_of_life

type coord struct{ i, j int }

func gameOfLife(board [][]int) {
	neighbors := func(i, j int) []coord {
		var res []coord
		for di := -1; di <= 1; di++ {
			for dj := -1; dj <= 1; dj++ {
				if di == 0 && dj == 0 {
					continue
				}
				if i+di < 0 || i+di >= len(board) {
					continue
				}
				if j+dj < 0 || j+dj >= len(board[i]) {
					continue
				}
				res = append(res, coord{i + di, j + dj})
			}
		}
		return res
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			live := 0
			for _, neigh := range neighbors(i, j) {
				if board[neigh.i][neigh.j]&0b1 == 1 {
					live++
				}
			}

			switch {
			case board[i][j]&0b1 == 0b1 && (live < 2 || live > 3):
				// do nothing
			case board[i][j]&0b1 == 0b1 || (board[i][j] == 0 && live == 3):
				board[i][j] |= 0b10
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j]&0b10 != 0 {
				board[i][j] = 1
			} else {
				board[i][j] = 0
			}
		}
	}
}
