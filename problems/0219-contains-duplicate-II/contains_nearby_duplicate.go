package _219_contains_duplicate_II

func containsNearbyDuplicate(nums []int, k int) bool {
	dict := map[int]int{}
	for i, n := range nums {
		if j, ok := dict[n]; ok && i-j <= k {
			return true
		}

		dict[n] = i
	}
	return false
}
