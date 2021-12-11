package _878_nth_magical_number

func nthMagicalNumber(n int, a int, b int) int {
	mod := 1000000007
	l := a / gcd(a, b) * b
	lo := 0
	hi := n * min(a, b)

	for lo < hi {
		mi := lo + (hi-lo)/2
		if mi/a+mi/b-mi/l < n {
			lo = mi + 1
		} else {
			hi = mi
		}
	}

	return lo % mod
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
