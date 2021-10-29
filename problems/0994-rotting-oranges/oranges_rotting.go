package _994_rotting_oranges

var directions = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

const (
	Fresh  = 1
	Rotten = 2
)

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var rotten [][2]int

	fresh := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch grid[i][j] {
			case Fresh:
				fresh++
			case Rotten:
				rotten = append(rotten, [2]int{i, j})
			}
		}
	}

	if fresh == 0 {
		return 0
	}

	minutes := 0
	for len(rotten) > 0 {
		for _, o := range rotten {
			for _, dir := range directions {
				x, y := o[0]+dir[0], o[1]+dir[1]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == Fresh {
					grid[x][y] = Rotten
					rotten = append(rotten, [2]int{x, y})
					fresh--
				}
			}
			rotten = rotten[1:]
		}

		minutes++
		if fresh == 0 {
			return minutes
		}
	}

	return -1
}
