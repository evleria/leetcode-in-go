package _790_domino_and_tromino_tiling

func numTilings(n int) int {
	if n == 1 {
		return 1
	}
	dp1 := make([]int, n+1)
	dp2 := make([]int, n+1)
	dp1[1], dp1[2] = 1, 2
	dp2[2] = 1
	for i := 3; i <= n; i++ {
		dp1[i] = (dp1[i-2] + dp1[i-1] + dp2[i-1]*2) % 1000000007
		dp2[i] = (dp1[i-2] + dp2[i-1]) % 1000000007
	}
	return dp1[n]
}
