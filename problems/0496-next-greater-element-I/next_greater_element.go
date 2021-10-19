package _496_next_greater_element_I

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))
	m := make(map[int]int, len(nums2))
	for i, ch := range nums2 {
		m[ch] = i
	}

	for i := 0; i < len(nums1); i++ {
		found := false
		for j := m[nums1[i]] + 1; j < len(nums2); j++ {
			if nums2[j] > nums1[i] {
				ans[i] = nums2[j]
				found = true
				break
			}
		}
		if !found {
			ans[i] = -1
		}
	}
	return ans
}
