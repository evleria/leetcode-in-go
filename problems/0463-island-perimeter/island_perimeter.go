package _463_island_perimeter

func islandPerimeter(grid [][]int) int {
	result := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				result += 4
				if i > 0 && grid[i-1][j] == 1 {
					result -= 2
				}
				if j > 0 && grid[i][j-1] == 1 {
					result -= 2
				}
			}
		}
	}
	return result
}
