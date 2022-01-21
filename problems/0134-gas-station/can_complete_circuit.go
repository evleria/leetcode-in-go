package _134_gas_station

func canCompleteCircuit(gas []int, cost []int) int {
	sum, max, j := 0, 0, 0
	for i := 0; i < len(cost); i++ {
		sum += gas[i] - cost[i]
		max += gas[i] - cost[i]
		if max < 0 {
			max = 0
			j = i + 1
		}
	}
	if sum >= 0 {
		return j
	}
	return -1
}
