package _306_jump_game_III

func canReach(arr []int, start int) bool {
	if start >= len(arr) || start < 0 {
		return false
	}

	val := arr[start]

	if val == 0 {
		return true
	}

	if val < 0 {
		return false
	}

	arr[start] *= -1

	return canReach(arr, start+val) || canReach(arr, start-val)
}
