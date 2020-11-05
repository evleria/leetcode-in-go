package _088_merge_sorted_array

func merge(nums1 []int, m int, nums2 []int, n int) {
	out, i1, i2 := m+n-1, m-1, n-1

	for i1 >= 0 && i2 >= 0 {
		if nums1[i1] > nums2[i2] {
			nums1[out] = nums1[i1]
			i1--
		} else {
			nums1[out] = nums2[i2]
			i2--
		}
		out--
	}

	for i2 >= 0 {
		nums1[out] = nums2[i2]
		i2--
		out--
	}
}
