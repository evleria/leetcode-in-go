package _154_keep_multiplying_found_values_by_two

func findFinalValue(nums []int, original int) int {
	m := map[int]bool{}
	for _, num := range nums {
		m[num] = true
	}
	for m[original] {
		original *= 2
	}

	return original
}
