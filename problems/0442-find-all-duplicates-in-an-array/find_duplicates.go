package _442_find_all_duplicates_in_an_array

func findDuplicates(nums []int) []int {
	result := make([]int, 0, len(nums))
	m := make(map[int]bool, len(nums))
	for _, num := range nums {
		if _, ok := m[num]; ok {
			result = append(result, num)
		} else {
			m[num] = true
		}
	}
	return result
}
