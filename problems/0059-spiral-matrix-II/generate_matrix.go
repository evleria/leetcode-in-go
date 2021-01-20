package _059_spiral_matrix_II

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	counter := 1
	top, right, bottom, left := 0, n-1, n-1, 0
	for counter <= n*n {
		for i := left; i <= right; i++ {
			result[top][i] = counter
			counter++
		}
		top++
		for i := top; i <= bottom; i++ {
			result[i][right] = counter
			counter++
		}
		right--
		if top <= bottom {
			for i := right; i >= left; i-- {
				result[bottom][i] = counter
				counter++
			}
			bottom--
			for i := bottom; i >= top; i-- {
				result[i][left] = counter
				counter++
			}
			left++
		}
	}

	return result
}
