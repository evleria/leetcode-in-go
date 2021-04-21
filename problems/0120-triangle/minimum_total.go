package _120_triangle

import "math"

func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			triangle[i][j] += int(math.Min(float64(triangle[i+1][j+1]), float64(triangle[i+1][j])))
		}
	}
	return triangle[0][0]
}
