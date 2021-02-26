package _376_wiggle_subsequence

func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	prevDiff, counter := nums[1]-nums[0], 1
	if prevDiff != 0 {
		counter++
	}

	for i := 2; i < len(nums); i++ {
		if diff := nums[i] - nums[i-1]; (diff > 0 && prevDiff <= 0) || (diff < 0 && prevDiff >= 0) {
			counter++
			prevDiff = diff
		}
	}

	return counter
}
