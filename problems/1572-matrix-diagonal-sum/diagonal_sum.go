package _572_matrix_diagonal_sum

func diagonalSum(mat [][]int) int {
	sum := 0
	n := len(mat)
	for i := 0; i < n; i++ {
		sum += mat[i][i]
		if j := n - i - 1; i != j {
			sum += mat[i][j]
		}
	}
	return sum
}
