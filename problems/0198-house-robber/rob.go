package _198_house_robber

func rob(nums []int) int {
	var a, b int
	for _, n := range nums {
		a, b = b, max(b, a+n)
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
