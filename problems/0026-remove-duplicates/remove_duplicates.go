package _026_remove_duplicates

func removeDuplicates(nums []int) int {
	l := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[l-1] {
			nums[l] = nums[i]
			l++
		}
	}

	return l
}
