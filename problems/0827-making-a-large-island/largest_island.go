package _827_making_a_large_island

type coordinate [2]int
type island map[coordinate]bool

var directions = [4]coordinate{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

const idOffset = 1 << 20

func largestIsland(grid [][]int) int {
	n := len(grid)
	id := 1
	max := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				sum := 1
				ids := map[int]bool{}
				for _, dir := range directions {
					x, y := i+dir[0], j+dir[1]
					if x >= 0 && x < n && y >= 0 && y < n && grid[x][y] != 0 {
						area, islandId := getArea(grid, x, y, &id)
						if _, ok := ids[islandId]; !ok {
							sum += area
							ids[islandId] = true
						}
					}
				}

				if sum > max {
					max = sum
				}
			}
		}
	}
	if max == 0 {
		return n * n
	}
	return max
}

func getArea(grid [][]int, i, j int, id *int) (int, int) {
	if r := grid[i][j]; r/idOffset > 0 {
		return r % idOffset, r / idOffset
	}

	n := len(grid)
	var visitAdjacent func(visited island, i, j int)
	visitAdjacent = func(visited island, i, j int) {
		visited[coordinate{i, j}] = true
		for _, dir := range directions {
			x, y := i+dir[0], j+dir[1]

			if x >= 0 && x < n && y >= 0 && y < n && grid[x][y] != 0 && !visited[coordinate{x, y}] {
				visitAdjacent(visited, x, y)
			}
		}
	}

	visited := island{}
	visitAdjacent(visited, i, j)
	area, islandId := len(visited), *id
	*id++
	r := area + islandId*idOffset
	for c := range visited {
		grid[c[0]][c[1]] = r
	}
	return area, islandId
}
