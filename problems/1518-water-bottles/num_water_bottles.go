package _518_water_bottles

func numWaterBottles(numBottles int, numExchange int) int {
	result, empty := numBottles, numBottles
	for empty >= numExchange {
		numBottles, empty = empty/numExchange, empty%numExchange
		numBottles, empty, result = 0, numBottles+empty, result+numBottles
	}
	return result
}
