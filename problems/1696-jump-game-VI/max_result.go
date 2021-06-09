package _696_jump_game_VI

func maxResult(nums []int, k int) int {
	n := len(nums)
	dp, q := make([]int, n), []int{0}
	dp[0] = nums[0]

	for i := 1; i < n; i++ {
		if len(q) > 0 && q[0] < i-k {
			q = q[1:]
		}

		dp[i] = dp[q[0]] + nums[i]

		for len(q) > 0 && dp[q[len(q)-1]] < dp[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	return dp[n-1]
}
