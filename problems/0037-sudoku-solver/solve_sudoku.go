package _037_sudoku_solver

func solveSudoku(board [][]byte) {
	helper(board, 0, 0)
}

func helper(board [][]byte, i, j int) bool {
	if i == 9 {
		return true
	}
	if j == 9 {
		return helper(board, i+1, 0)
	}
	if board[i][j] != '.' {
		return helper(board, i, j+1)
	}
	for c := byte('1'); c <= '9'; c++ {
		if !isValid(board, i, j, c) {
			continue
		}
		board[i][j] = c
		if helper(board, i, j+1) {
			return true
		}
		board[i][j] = '.'
	}
	return false
}

func isValid(board [][]byte, i, j int, c byte) bool {
	for x := 0; x < 9; x++ {
		if board[x][j] == c {
			return false
		}
	}
	for y := 0; y < 9; y++ {
		if board[i][y] == c {
			return false
		}
	}
	xBase, yBase := i-i%3, j-j%3
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if board[xBase+x][yBase+y] == c {
				return false
			}
		}
	}
	return true
}
