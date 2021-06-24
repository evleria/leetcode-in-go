package _576_out_of_boundary_paths

var directions = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
var mod = 1_000_000_007

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	if maxMove == 0 {
		return 0
	}

	field := createField(m, n)
	field[startRow][startColumn] = 1

	count := 0
	for move := 0; move < maxMove; move++ {
		nextField := createField(m, n)

		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if s := field[i][j]; s > 0 {
					for _, dir := range directions {
						x, y := i+dir[0], j+dir[1]
						if x < 0 || x >= m || y < 0 || y >= n {
							count = (count + s) % mod
						} else {
							nextField[x][y] = (nextField[x][y] + s) % mod
						}
					}
				}
			}
		}

		field = nextField
	}

	return count
}

func createField(m, n int) [][]int {
	field := make([][]int, m)
	for i := range field {
		field[i] = make([]int, n)
	}
	return field
}
