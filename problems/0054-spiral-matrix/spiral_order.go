package _054_spiral_matrix

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0, m*n)

	top, right, bottom, left := 0, n-1, m-1, 0
	for len(result) < m*n {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--
		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
			if right >= left {
				for i := bottom; i >= top; i-- {
					result = append(result, matrix[i][left])
				}
				left++
			}
		}

	}
	return result
}
