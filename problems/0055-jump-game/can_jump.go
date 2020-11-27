package _055_jump_game

func canJump(nums []int) bool {
	for max, i := 0, 0; i <= max; i++ {
		if m := nums[i] + i; m > max {
			max = m
		}
		if max >= len(nums)-1 {
			return true
		}
	}

	return false
}
