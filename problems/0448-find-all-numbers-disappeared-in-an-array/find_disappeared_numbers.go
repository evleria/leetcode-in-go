package _448_find_all_numbers_disappeared_in_an_array

func findDisappearedNumbers(nums []int) []int {
	result := make([]int, len(nums))
	for _, n := range nums {
		result[n-1] = n
	}

	out := 0
	for i := 0; i < len(result); i++ {
		if result[i] == 0 {
			result[out] = i + 1
			out++
		}
	}

	return result[:out]
}
