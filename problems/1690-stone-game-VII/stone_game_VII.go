package _690_stone_game_VII

func stoneGameVII(stones []int) int {
	n := len(stones)
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, n)
	}
	preSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + stones[i]
	}
	getSum := func(left, right int) int {
		return preSum[right+1] - preSum[left]
	}
	var dp func(int, int) int
	dp = func(left, right int) int {
		if left == right {
			return 0
		}
		if r := memo[left][right]; r != 0 {
			return r
		}
		scoreRemoveLeftMost := getSum(left+1, right)
		scoreRemoveRightMost := getSum(left, right-1)
		ans := scoreRemoveLeftMost - dp(left+1, right)
		if rightAns := scoreRemoveRightMost - dp(left, right-1); rightAns > ans {
			ans = rightAns
		}
		memo[left][right] = ans
		return ans
	}
	return dp(0, n-1)
}
