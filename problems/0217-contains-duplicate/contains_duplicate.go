package _217_contains_duplicate

func containsDuplicate(nums []int) bool {
	dict := map[int]bool{}
	for _, n := range nums {
		if dict[n] {
			return true
		}

		dict[n] = true
	}
	return false
}
