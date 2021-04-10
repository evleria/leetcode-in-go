package _329_longest_increasing_path_in_a_matrix

var directions = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

// m - vertical, 1st dim (x)
// n - horizontal, 2nd dim (y)
func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
	}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if r := memo[i][j]; r != 0 {
			return r
		}

		max := 0
		for _, dir := range directions {
			x, y := i+dir[0], j+dir[1]
			if x >= 0 && x < m && y >= 0 && y < n && matrix[x][y] < matrix[i][j] {
				if r := dfs(x, y); r > max {
					max = r
				}
			}
		}

		r := max + 1
		memo[i][j] = r
		return r
	}

	max := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if r := dfs(i, j); r > max {
				max = r
			}
		}
	}
	return max
}
