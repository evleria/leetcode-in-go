package _275_find_winner_on_a_tic_tac_toe_game

func tictactoe(moves [][]int) string {
	board := [3][3]rune{}

	for i, move := range moves {
		ch := 'A'
		if i%2 != 0 {
			ch = 'B'
		}
		board[move[0]][move[1]] = ch
	}

	for i := 0; i < 3; i++ {
		chH, chV := board[i][0], board[0][i]
		if chH != 0 && board[i][1] == chH && board[i][2] == chH {
			return string(chH)
		}
		if chV != 0 && board[1][i] == chV && board[2][i] == chV {
			return string(chV)
		}
	}

	ch := board[1][1]
	if ch != 0 && board[0][0] == ch && board[2][2] == ch {
		return string(ch)
	}
	if ch != 0 && board[0][2] == ch && board[2][0] == ch {
		return string(ch)
	}

	if len(moves) == 9 {
		return "Draw"
	}

	return "Pending"
}
