package _665_non_decreasing_array

func checkPossibility(nums []int) bool {
	modified := false
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if modified {
				return false
			}
			if i == 0 || nums[i-1] <= nums[i+1] {
				nums[i] = nums[i+1]
			} else {
				nums[i+1] = nums[i]
			}
			modified = true
		}
	}
	return true
}
