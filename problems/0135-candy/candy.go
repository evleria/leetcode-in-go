package _135_candy

import "math"

func candy(ratings []int) int {
	if len(ratings) == 1 {
		return 1
	}

	candies, up, down, oldSlope := 0, 0, 0, 0
	for i := 1; i < len(ratings); i++ {
		newSlope := getSign(ratings[i] - ratings[i-1])

		if (oldSlope > 0 && newSlope == 0) || (oldSlope < 0 && newSlope >= 0) {
			candies += calcSum(up) + calcSum(down) + int(math.Max(float64(up), float64(down)))
			up, down = 0, 0
		}

		switch newSlope {
		case 1:
			up++
		case -1:
			down++
		default:
			candies++
		}
		oldSlope = newSlope
	}
	candies += calcSum(up) + calcSum(down) + int(math.Max(float64(up), float64(down))) + 1
	return candies
}

func getSign(n int) int {
	switch {
	case n > 0:
		return 1
	case n < 0:
		return -1
	default:
		return 0
	}
}

func calcSum(n int) int {
	return ((n + 1) * n) / 2
}
