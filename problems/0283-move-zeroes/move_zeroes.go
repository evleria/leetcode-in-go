package _283_move_zeroes

func moveZeroes(nums []int) {
	out := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[out], nums[i] = nums[i], nums[out]
			out++
		}
	}
}
