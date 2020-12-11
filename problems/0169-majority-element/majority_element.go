package _169_majority_element

func majorityElement(nums []int) int {
	candidate, rating := nums[0], 1

	for i := 1; i < len(nums); i++ {
		if rating == 0 {
			candidate, rating = nums[i], 1
		} else {
			if nums[i] == candidate {
				rating++
			} else {
				rating--
			}
		}
	}

	return candidate
}
