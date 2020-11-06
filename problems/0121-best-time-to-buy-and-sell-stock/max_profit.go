package _121_best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice, maxProfit := prices[0], 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if profit := prices[i] - minPrice; profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}
