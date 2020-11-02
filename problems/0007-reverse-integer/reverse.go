package _007_reverse_integer

import "math"

func reverse(x int) int {
	abs, sign := abs(x)

	m := 0
	for ; abs > 0; abs /= 10 {
		m *= 10
		m += abs % 10
	}

	if result := m * sign; result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	} else {
		return result
	}
}

func abs(x int) (int, int) {
	if x >= 0 {
		return x, 1
	}
	return -x, -1
}
