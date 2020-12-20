package _034_find_first_and_last_position_of_element_in_sorted_array

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start + end) / 2
		switch {
		case target < nums[mid]:
			end = mid - 1
		case target > nums[mid]:
			start = mid + 1
		default:
			for start = mid; start > 0 && nums[start-1] == target; start-- {
			}
			for end = mid; end < len(nums)-1 && nums[end+1] == target; end++ {
			}

			return []int{start, end}
		}
	}

	return []int{-1, -1}
}
