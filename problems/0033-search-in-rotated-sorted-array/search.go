package _033_search_in_rotated_sorted_array

func search(nums []int, target int) int {
	n := len(nums)

	left, right := 0, n-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	pivot := left

	left, right = pivot, pivot-1+n
	for left <= right {
		mid := (left + right) / 2
		pivotedMid := mid % n
		switch {
		case nums[pivotedMid] > target:
			right = mid - 1
		case nums[pivotedMid] < target:
			left = mid + 1
		default:
			return pivotedMid
		}
	}

	return -1
}
