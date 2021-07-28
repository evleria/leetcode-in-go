package _252_cells_with_odd_values_in_a_matrix

func oddCells(m int, n int, indices [][]int) int {
	rows, cols := make([]int, m), make([]int, n)
	for _, index := range indices {
		rows[index[0]]++
		cols[index[1]]++
	}
	result := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result += (rows[i] + cols[j]) % 2
		}
	}
	return result
}
