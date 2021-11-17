package _961_n_repeated_element_in_size_2n_array

func repeatedNTimes(nums []int) int {
	seen := make(map[int]bool, len(nums))
	for _, num := range nums {
		if seen[num] {
			return num
		}
		seen[num] = true
	}
	return 0
}
