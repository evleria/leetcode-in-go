package _054_spiral_matrix

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0, m*n)
	loops := (min(m, n) + 1) / 2

	for x := 0; x < loops; x++ {
		for j := x; j < n-x; j++ {
			result = append(result, matrix[x][j])
		}
		for i := 1 + x; i < m-1-x; i++ {
			result = append(result, matrix[i][n-1-x])
		}
		if x != m-1-x {
			for j := n - 1 - x; j >= x; j-- {
				result = append(result, matrix[m-1-x][j])
			}
			if x != n-1-x {
				for i := m - 2 - x; i >= 1+x; i-- {
					result = append(result, matrix[i][x])
				}
			}
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
