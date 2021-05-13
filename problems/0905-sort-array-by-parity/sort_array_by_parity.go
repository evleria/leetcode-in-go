package _905_sort_array_by_parity

func sortArrayByParity(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		switch {
		case nums[left]%2 == 0:
			left++
		case nums[right]%2 != 0:
			right--
		default:
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}

	return nums
}
