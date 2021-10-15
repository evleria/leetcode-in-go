package _309_best_time_to_buy_and_sell_stock_with_cooldown

func maxProfit(prices []int) int {
	if prices == nil || len(prices) <= 1 {
		return 0
	}

	x := -prices[0]
	y, a, b, c := x, 0, 0, 0

	for i := 0; i < len(prices); i++ {
		x = max(y, c-prices[i])
		a = max(b, y+prices[i])
		y, c, b = x, b, a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
