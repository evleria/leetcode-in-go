package _906_super_palindromes

import (
	"math"
	"strconv"
)

func superpalindromesInRange(left string, right string) int {
	result := 0
	l, r := sqrt(left, math.Ceil), sqrt(right, math.Floor)

	countSuperpalindromes := func(isEven bool) {
		for i := 1; ; i++ {
			p, j := i, i
			if !isEven {
				j /= 10
			}
			for ; j > 0; j /= 10 {
				p *= 10
				p += j % 10
			}
			if p > r {
				break
			}
			if p >= l && isPalindrome(p*p) {
				result++
			}
		}
	}

	countSuperpalindromes(true)
	countSuperpalindromes(false)

	return result
}

func isPalindrome(x int) bool {
	y := 0
	for ; x > y; x /= 10 {
		y *= 10
		y += x % 10
	}

	return x == y || x == y/10
}

func sqrt(s string, roundFunc func(float64) float64) int {
	f, _ := strconv.ParseFloat(s, 64)
	return int(roundFunc(math.Sqrt(f)))
}
