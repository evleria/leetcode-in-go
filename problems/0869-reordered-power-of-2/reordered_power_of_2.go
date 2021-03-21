package _869_reordered_power_of_2

import "math"

func reorderedPowerOf2(N int) bool {
	c := counter(N)
	for i := 1; i <= 1_000_000_000; i <<= 1 {
		if counter(i) == c {
			return true
		}
	}
	return false
}

func counter(n int) float64 {
	c := 0.0
	for ; n > 0; n /= 10 {
		c += math.Pow10(n % 10)
	}
	return c
}
