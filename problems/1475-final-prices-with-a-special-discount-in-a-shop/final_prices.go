package _475_final_prices_with_a_special_discount_in_a_shop

func finalPrices(prices []int) []int {
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			s := prices[j]
			if s <= prices[i] {
				prices[i] -= s
				break
			}
		}
	}
	return prices
}
