package _081_search_in_rotated_sorted_array_II

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	l := 0
	r := len(nums) - 1

	for l+1 < r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		}
		if nums[l] < nums[mid] {
			if nums[l] <= target && nums[mid] > target {
				r = mid
			} else {
				l = mid
			}
		} else if nums[l] > nums[mid] {
			if nums[mid] < target && nums[r] >= target {
				l = mid
			} else {
				r = mid
			}
		} else {
			l++
		}
	}
	if nums[l] == target || nums[r] == target {
		return true
	} else {
		return false
	}
}
