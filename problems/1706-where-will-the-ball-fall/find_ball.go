package _706_where_will_the_ball_fall

func findBall(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	result := make([]int, n)

	for i := 0; i < n; i++ {
		pos := i
		for j := 0; j < m; j++ {
			dir := grid[j][pos]
			if nextPos := pos + dir; nextPos < 0 || nextPos >= n || grid[j][nextPos] == -dir {
				pos = -1
				break
			}
			pos += dir
		}
		result[i] = pos
	}

	return result
}
