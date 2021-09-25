package _293_shortest_path_in_a_grid_with_obstacles_elimination

var moves = [][]int{
	{1, 0},
	{0, 1},
	{0, -1},
	{-1, 0},
}

func shortestPath(grid [][]int, k int) int {
	stack := make([][]int, 0)
	visited := make([][]int, len(grid))
	for i := range visited {
		visited[i] = make([]int, len(grid[0]))
		for j := range visited[i] {
			visited[i][j] = -1
		}
	}
	visited[0][0] = k
	stack = append(stack, []int{0, 0, k})
	level := 0
	for len(stack) != 0 {
		size := len(stack)
		for i := range stack {
			if stack[i][0] == len(grid)-1 && stack[i][1] == len(grid[0])-1 {
				return level
			}
			for _, move := range moves {
				nextx := stack[i][0] + move[0]
				nexty := stack[i][1] + move[1]
				if nextx >= 0 && nextx < len(grid) && nexty >= 0 && nexty < len(grid[0]) {
					curK := stack[i][2]
					if grid[nextx][nexty] == 1 {
						curK--
					}
					if curK >= 0 {
						if visited[nextx][nexty] == -1 || (visited[nextx][nexty] != -1 && curK > visited[nextx][nexty]) {
							visited[nextx][nexty] = curK
							stack = append(stack, []int{nextx, nexty, curK})
						}
					}
				}
			}
		}
		stack = stack[size:]
		level++
	}
	return -1
}
