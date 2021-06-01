package _695_max_area_of_island

var directions = [][2]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				if area := dfs(i, j, grid); area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

func dfs(x, y int, grid [][]int) int {
	grid[x][y] = 0
	area := 1
	for _, direction := range directions {
		x1, y1 := x+direction[0], y+direction[1]

		if x1 >= 0 && x1 < len(grid) && y1 >= 0 && y1 < len(grid[0]) && grid[x1][y1] == 1 {
			area += dfs(x1, y1, grid)
		}
	}
	return area
}
