package _122_best_time_to_buy_and_sell_stock_II

func maxProfit(prices []int) int {
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		if delta := prices[i+1] - prices[i]; delta > 0 {
			profit += delta
		}
	}
	return profit
}
