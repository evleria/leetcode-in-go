package _045_jump_game_II

func jump(nums []int) int {
	curJump, farthestJump, jumps := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > farthestJump {
			farthestJump = i + nums[i]
		}

		if i == curJump {
			jumps, curJump = jumps+1, farthestJump

			if curJump >= len(nums)-1 {
				return jumps
			}
		}
	}

	return 0
}
