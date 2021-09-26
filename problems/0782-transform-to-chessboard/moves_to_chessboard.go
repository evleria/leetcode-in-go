package _782_transform_to_chessboard

func movesToChessboard(board [][]int) int {
	n, rowSum, colSum, rowSwap, colSwap := len(board), 0, 0, 0, 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[0][0]^board[i][0]^board[0][j]^board[i][j] == 1 {
				return -1
			}
		}
	}

	for i := 0; i < n; i++ {
		rowSum += board[0][i]
		colSum += board[i][0]
		if board[i][0] == i%2 {
			rowSwap++
		}
		if board[0][i] == i%2 {
			colSwap++
		}
	}

	if rowSum != n/2 && rowSum != (n+1)/2 {
		return -1
	}
	if colSum != n/2 && colSum != (n+1)/2 {
		return -1
	}
	if n%2 == 1 {
		if colSwap%2 == 1 {
			colSwap = n - colSwap
		}
		if rowSwap%2 == 1 {
			rowSwap = n - rowSwap
		}
	} else {
		colSwap = min(n-colSwap, colSwap)
		rowSwap = min(n-rowSwap, rowSwap)
	}
	return (colSwap + rowSwap) / 2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
