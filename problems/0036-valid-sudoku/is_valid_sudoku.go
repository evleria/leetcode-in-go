package _036_valid_sudoku

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		var digits int
		for j := 0; j < 9; j++ {
			if checkAndSet(board[i][j], &digits) {
				return false
			}
		}
	}

	for i := 0; i < 9; i++ {
		var digits int
		for j := 0; j < 9; j++ {
			if checkAndSet(board[j][i], &digits) {
				return false
			}
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			iOff, jOff := i*3, j*3
			var digits int
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					if checkAndSet(board[x+iOff][y+jOff], &digits) {
						return false
					}
				}
			}
		}
	}

	return true
}

func checkAndSet(ch byte, digits *int) bool {
	if ch != '.' {
		bit := 1 << (ch - '0')
		if *digits&bit != 0 {
			return true
		}
		*digits |= bit
	}
	return false
}
