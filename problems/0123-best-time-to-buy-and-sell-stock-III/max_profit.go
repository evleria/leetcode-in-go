package _123_best_time_to_buy_and_sell_stock_III

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	start := make([]int, len(prices))
	minPrice, maxProfit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if profit := prices[i] - minPrice; profit > maxProfit {
			maxProfit = profit
		}
		start[i] = maxProfit
	}

	if len(prices) < 4 {
		return maxProfit
	}

	result := maxProfit
	end := make([]int, len(prices))
	maxPrice, maxProfit := prices[len(prices)-1], 0
	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i] > maxPrice {
			maxPrice = prices[i]
		} else if profit := maxPrice - prices[i]; profit > maxProfit {
			maxProfit = profit
		}
		end[i] = maxProfit
	}

	for i := 1; i < len(prices)-2; i++ {
		if profit := start[i] + end[i+1]; profit > result {
			result = profit
		}
	}

	return result
}
