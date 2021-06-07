package _746_min_cost_climbing_stairs

func minCostClimbingStairs(cost []int) int {
	a, b := 0, 0
	for i := 2; i <= len(cost); i++ {
		a, b = b, min(cost[i-2]+a, cost[i-1]+b)
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
