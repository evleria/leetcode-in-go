package _200_number_of_islands

type coordinate struct {
	x int
	y int
}

var directions = []coordinate{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}

func numIslands(grid [][]byte) int {
	counter, visited := 0, make([][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[i]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				counter++

				stack := make([]coordinate, 1, 16)
				stack[0] = coordinate{i, j}

				for len(stack) != 0 {
					cur := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					x, y := cur.x, cur.y
					visited[x][y] = true

					for _, direction := range directions {
						x1, y1 := x+direction.x, y+direction.y

						if x1 >= 0 && x1 < len(grid) && y1 >= 0 && y1 < len(grid[0]) && !visited[x1][y1] && grid[x1][y1] == '1' {
							stack = append(stack, coordinate{x1, y1})
						}
					}
				}
			}
		}
	}

	return counter
}
