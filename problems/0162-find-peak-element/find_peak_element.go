package _162_find_peak_element

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		m := (left + right) / 2
		if nums[m] > nums[m+1] {
			right = m
		} else {
			left = m + 1
		}
	}
	return left
}
