package _867_transpose_matrix

func transpose(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, m)
		for j := 0; j < m; j++ {
			result[i][j] = matrix[j][i]
		}
	}
	return result
}
