package _342_power_of_four

func isPowerOfFour(n int) bool {
	return n&(n-1) == 0 && n%3 == 1
}
