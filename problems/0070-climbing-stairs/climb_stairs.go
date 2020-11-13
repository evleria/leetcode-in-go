package _070_climbing_stairs

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	a, b := 2, 3
	for i := 5; i <= n; i++ {
		a, b = b, a+b
	}
	return a + b
}
