package _718_maximum_length_of_repeated_subarray

func findLength(nums1 []int, nums2 []int) int {
	m := map[int][]int{}
	for i, n := range nums2 {
		m[n] = append(m[n], i)
	}

	max := 0
	for i := 0; i < len(nums1)-max; i++ {
		for _, j := range m[nums1[i]] {
			offset := 1
			for ; i+offset < len(nums1) && j+offset < len(nums2) && nums1[i+offset] == nums2[j+offset]; offset++ {
			}

			if offset > max {
				max = offset
			}
		}
	}
	return max
}
