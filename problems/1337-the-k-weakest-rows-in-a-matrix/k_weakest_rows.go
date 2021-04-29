package _337_the_k_weakest_rows_in_a_matrix

func kWeakestRows(mat [][]int, k int) []int {
	m, n := len(mat), len(mat[0])
	result, taken := make([]int, 0, k), make([]bool, m)
	for i := 0; i <= n; i++ {
		for j := 0; j < m; j++ {
			if !taken[j] && (i == n || mat[j][i] == 0) {
				result, taken[j] = append(result, j), true
				if len(result) == k {
					return result
				}
			}
		}
	}

	return nil
}
