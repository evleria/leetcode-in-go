package _365_how_many_numbers_are_smaller_than_the_current_number

func smallerNumbersThanCurrent(nums []int) []int {
	counters := [101]int{}
	for _, n := range nums {
		counters[n]++
	}

	prev := counters[0]
	counters[0] = 0
	for i := 1; i <= 100; i++ {
		counters[i], prev = prev, prev+counters[i]
	}

	result := make([]int, len(nums))
	for i, n := range nums {
		result[i] = counters[n]
	}

	return result
}
