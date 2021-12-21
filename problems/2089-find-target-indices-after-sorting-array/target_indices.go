package _089_find_target_indices_after_sorting_array

func targetIndices(nums []int, target int) []int {
	less, count := 0, 0
	for _, num := range nums {
		if num < target {
			less++
		} else if num == target {
			count++
		}
	}

	result := make([]int, count)

	for i := 0; i < count; i++ {
		result[i] = less + i
	}

	return result
}
