package _118_pascals_triangle

func generate(numRows int) [][]int {

	result := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		switch i {
		case 0:
			result[i] = []int{1}
		case 1:
			result[i] = []int{1, 1}
		default:
			result[i] = make([]int, i+1)
			result[i][0], result[i][i] = 1, 1
			for j := 1; j < i; j++ {
				result[i][j] = result[i-1][j-1] + result[i-1][j]
			}
		}
	}

	return result
}
