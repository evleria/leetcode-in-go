package _074_number_of_submatrices_that_sum_to_target

// x, horizontal, i, first,  n
// y, vertical,   j, second, m
func numSubmatrixSumTarget(matrix [][]int, target int) int {
	n, m := len(matrix), len(matrix[0])
	sums := make([][]int, n)
	for i := 0; i < n; i++ {
		sums[i] = make([]int, m)
	}

	sums[0][0] = matrix[0][0]
	for i := 1; i < n; i++ {
		sums[i][0] = sums[i-1][0] + matrix[i][0]
	}
	for j := 1; j < m; j++ {
		sums[0][j] = sums[0][j-1] + matrix[0][j]
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			sums[i][j] = sums[i-1][j] + sums[i][j-1] - sums[i-1][j-1] + matrix[i][j]
		}
	}

	result := 0
	for l := 0; l < n; l++ {
		for r := l; r < n; r++ {
			for t := 0; t < m; t++ {
				for b := t; b < m; b++ {
					sum := sums[r][b]
					if l > 0 {
						sum -= sums[l-1][b]
					}
					if t > 0 {
						sum -= sums[r][t-1]
					}
					if l > 0 && t > 0 {
						sum += sums[l-1][t-1]
					}

					if sum == target {
						result++
					}
				}
			}
		}
	}
	return result
}
