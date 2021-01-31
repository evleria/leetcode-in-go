package _566_reshape_the_matrix

func matrixReshape(nums [][]int, r int, c int) [][]int {
	n, m := len(nums), len(nums[0])
	if r*c != n*m {
		return nums
	}
	result := make([][]int, r)
	for i := 0; i < r; i++ {
		result[i] = make([]int, c)
		for j := 0; j < c; j++ {
			x := i*c + j
			result[i][j] = nums[x/m][x%m]
		}
	}
	return result
}
