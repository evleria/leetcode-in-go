package _204_count_primes

import "math"

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	eliminated, end := make([]bool, n), int(math.Sqrt(float64(n)))
	for i := 3; i <= end; i += 2 {
		if !eliminated[i] {
			for j := i * i; j < n; j += i {
				eliminated[j] = true
			}
		}
	}

	count := 1
	for i := 3; i < n; i += 2 {
		if !eliminated[i] {
			count++
		}
	}

	return count
}
