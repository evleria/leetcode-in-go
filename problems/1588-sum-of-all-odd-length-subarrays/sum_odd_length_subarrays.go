package _588_sum_of_all_odd_length_subarrays

func sumOddLengthSubarrays(arr []int) int {
	result, l := 0, len(arr)
	for i, n := range arr {
		result += ((i+1)*(l-i) + 1) / 2 * n
	}
	return result
}
