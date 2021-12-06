package _217_minimum_cost_to_move_chips_to_the_same_position

func minCostToMoveChips(position []int) int {
	counters := [2]int{} // [even, odd]
	for _, p := range position {
		counters[p%2]++
	}
	if counters[0] < counters[1] {
		return counters[0]
	}
	return counters[1]
}
