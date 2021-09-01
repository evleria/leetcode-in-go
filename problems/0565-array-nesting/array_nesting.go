package _565_array_nesting

func arrayNesting(nums []int) int {
	maxSize := 0
	for i := 0; i < len(nums); i++ {
		size := 0
		for j := i; nums[j] >= 0; size++ {
			nums[j], j = -1, nums[j]
		}
		if size > maxSize {
			maxSize = size
		}
	}
	return maxSize
}
