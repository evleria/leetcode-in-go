package _714_best_time_to_buy_and_sell_stock_with_transaction_fee

func maxProfit(prices []int, fee int) int {
	days := len(prices)
	buy, sell := -prices[0]-fee, 0
	for i := 1; i < days; i++ {
		buy, sell = max(buy, sell-prices[i]-fee), max(sell, buy+prices[i])
	}
	return sell
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
