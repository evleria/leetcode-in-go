package _349_intersection_of_two_arrays

func intersection(nums1 []int, nums2 []int) []int {
	result, seen := make([]int, 0, 16), make(map[int]bool, 16)
	for _, n := range nums1 {
		seen[n] = true
	}
	for _, n := range nums2 {
		if _, ok := seen[n]; ok {
			result = append(result, n)
			delete(seen, n)
		}
	}
	return result
}
