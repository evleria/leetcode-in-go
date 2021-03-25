package _417_pacific_atlantic_water_flow

var directions = [][2]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}

func pacificAtlantic(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return [][]int{}
	}

	m, n := len(matrix[0]), len(matrix)

	visited := make([][][2]bool, n)
	pacific := make([][2]int, 0, m+n-1)
	atlantic := make([][2]int, 0, m+n-1)

	for i := 0; i < n; i++ {
		visited[i] = make([][2]bool, m)
		pacific = append(pacific, [2]int{i, 0})
		atlantic = append(atlantic, [2]int{i, m - 1})
	}
	for i := 0; i < m; i++ {
		pacific = append(pacific, [2]int{0, i})
		atlantic = append(atlantic, [2]int{n - 1, i})
	}

	expand := func(ocean [][2]int, shore int) {
		for len(ocean) > 0 {
			c := ocean[len(ocean)-1]
			ocean = ocean[:len(ocean)-1]

			visited[c[0]][c[1]][shore] = true

			for _, dir := range directions {
				x := [2]int{c[0] + dir[0], c[1] + dir[1]}
				if x[0] >= 0 && x[0] < n && x[1] >= 0 && x[1] < m && !visited[x[0]][x[1]][shore] && matrix[x[0]][x[1]] >= matrix[c[0]][c[1]] {
					ocean = append(ocean, x)
				}
			}
		}
	}

	expand(pacific, 0)
	expand(atlantic, 1)

	var result [][]int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if visited[i][j] == [2]bool{true, true} {
				result = append(result, []int{i, j})
			}
		}
	}
	return result
}
