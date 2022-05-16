package _091_shortest_path_in_binary_matrix

var dirs = [8][2]int8{{-1, 1}, {0, 1}, {1, 1}, {-1, 0}, {1, 0}, {-1, -1}, {0, -1}, {1, -1}}

func shortestPathBinaryMatrix(grid [][]int) int {
	var n, m = int8(len(grid)), int8(len(grid[0]))
	if grid[0][0] == 1 || grid[n-1][m-1] == 1 {
		return -1
	}

	queue := [][2]int8{{0, 0}}
	grid[0][0] = 1

	pathLen := 0
	for len(queue) > 0 {
		pathLen++
		for _, c := range queue {
			if c[0] == n-1 && c[1] == m-1 {
				return pathLen
			}

			for _, d := range dirs {
				x, y := c[0]+d[0], c[1]+d[1]
				if x >= 0 && x < n && y >= 0 && y < m && grid[x][y] == 0 {
					grid[x][y] = 1
					queue = append(queue, [2]int8{x, y})
				}
			}
			queue = queue[1:]
		}
	}

	return -1
}
