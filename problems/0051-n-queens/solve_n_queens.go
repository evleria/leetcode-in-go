package _051_n_queens

func solveNQueens(n int) [][]string {
	result := [][]string{}
	backtrack(NewBoard(n), n, 0, &result)
	return result
}

func backtrack(board *Board, left int, start int, result *[][]string) {
	if left == 0 {
		*result = append(*result, board.String())
		return
	}

	for i := start; i < board.Size(); i++ {
		for j := 0; j < board.Size(); j++ {
			if board.CanSet(i, j) {
				board.Set(i, j)
				backtrack(board, left-1, i+1, result)
				board.Unset(i, j)
			}
		}
	}
}

type Board struct {
	board       [][]bool
	horizontals []bool
	verticals   []bool
	tlDiagonal  []bool
	trDiagonal  []bool
}

func NewBoard(n int) *Board {
	board := make([][]bool, n)
	for i := 0; i < n; i++ {
		board[i] = make([]bool, n)
	}

	return &Board{
		board:       board,
		horizontals: make([]bool, n),
		verticals:   make([]bool, n),
		tlDiagonal:  make([]bool, 2*n-1),
		trDiagonal:  make([]bool, 2*n-1),
	}
}

func (b *Board) String() []string {
	result := make([]string, b.Size())
	line := make([]byte, b.Size())
	for i, row := range b.board {
		for j, v := range row {
			if v {
				line[j] = 'Q'
			} else {
				line[j] = '.'
			}
		}
		result[i] = string(line)
	}

	return result
}

func (b *Board) CanSet(i, j int) bool {
	return !b.horizontals[i] && !b.verticals[j] && !b.tlDiagonal[i-j+b.Size()-1] && !b.trDiagonal[i+j]
}

func (b *Board) Set(i, j int) {
	b.board[i][j], b.horizontals[i], b.verticals[j], b.tlDiagonal[i-j+b.Size()-1], b.trDiagonal[i+j] = true, true, true, true, true
}

func (b *Board) Unset(i, j int) {
	b.board[i][j], b.horizontals[i], b.verticals[j], b.tlDiagonal[i-j+b.Size()-1], b.trDiagonal[i+j] = false, false, false, false, false
}

func (b *Board) Size() int {
	return len(b.board)
}
