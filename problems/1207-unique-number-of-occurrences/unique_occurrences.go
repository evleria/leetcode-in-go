package _207_unique_number_of_occurrences

func uniqueOccurrences(arr []int) bool {
	counters := map[int]int{}
	for _, n := range arr {
		counters[n]++
	}
	occurrences := map[int]bool{}
	for _, n := range counters {
		if occurrences[n] {
			return false
		}
		occurrences[n] = true
	}
	return true
}
