package _079_word_search

var directions = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func exist(board [][]byte, word string) bool {
	bytes := []byte(word)
	for i, row := range board {
		for j, ch := range row {
			if word[0] == ch {
				if dfs(board, [2]int{i, j}, bytes[1:]) {
					return true
				}
			}
		}
	}
	return false
}

func dfs(board [][]byte, src [2]int, rest []byte) bool {
	if len(rest) == 0 {
		return true
	}

	cur := board[src[0]][src[1]]
	board[src[0]][src[1]] = '.'
	for _, dir := range directions {
		dest := [2]int{src[0] + dir[0], src[1] + dir[1]}
		if dest[0] < len(board) && dest[0] >= 0 && dest[1] < len(board[0]) && dest[1] >= 0 && board[dest[0]][dest[1]] == rest[0] {
			if dfs(board, dest, rest[1:]) {
				return true
			}
		}
	}
	board[src[0]][src[1]] = cur

	return false
}
