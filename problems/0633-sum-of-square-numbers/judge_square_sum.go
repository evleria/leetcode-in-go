package _633_sum_of_square_numbers

import "math"

func judgeSquareSum(c int) bool {
	for i := 0; i*i <= c; i++ {
		s := math.Sqrt(float64(c - i*i))
		if float64(int(s)) == s {
			return true
		}
	}
	return false
}
