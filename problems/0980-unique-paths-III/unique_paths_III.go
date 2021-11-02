package _980_unique_paths_III

var directions = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

const (
	Empty    = 0
	Start    = 1
	Exit     = 2
	Obstacle = -1
)

func uniquePathsIII(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	var x, y int
	empty := 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			p := grid[i][j]
			if p == Empty {
				empty++
			} else if p == Start {
				x, y = i, j
			}
		}
	}

	ways := 0
	var dfs func(x, y int)
	dfs = func(x, y int) {
		grid[x][y] = Obstacle
		empty--

		for _, dir := range directions {
			i, j := x+dir[0], y+dir[1]

			if i >= 0 && i < m && j >= 0 && j < n {
				switch grid[i][j] {
				case Exit:
					if empty == 0 {
						ways++
					}
				case Empty:
					dfs(i, j)
				}
			}
		}

		grid[x][y] = Empty
		empty++
	}

	dfs(x, y)
	return ways
}
