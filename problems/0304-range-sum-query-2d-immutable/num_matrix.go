package _304_range_sum_query_2d_immutable

type NumMatrix struct {
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	return NumMatrix{
		sums: buildPrefixSums(matrix),
	}
}

func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	result := nm.sums[row2][col2]

	if col1 > 0 {
		result -= nm.sums[row2][col1-1]
	}
	if row1 > 0 {
		result -= nm.sums[row1-1][col2]
	}
	if row1 > 0 && col1 > 0 {
		result += nm.sums[row1-1][col1-1]
	}

	return result
}

func buildPrefixSums(matrix [][]int) [][]int {
	n, m := len(matrix), len(matrix[0])
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, m)
	}

	result[0][0] = matrix[0][0]
	for i := 1; i < m; i++ {
		result[0][i] = result[0][i-1] + matrix[0][i]
	}
	for i := 1; i < n; i++ {
		result[i][0] = result[i-1][0] + matrix[i][0]
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			result[i][j] = result[i-1][j] + result[i][j-1] + matrix[i][j] - result[i-1][j-1]
		}
	}

	return result
}
