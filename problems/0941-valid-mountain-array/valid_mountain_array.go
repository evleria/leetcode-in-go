package _941_valid_mountain_array

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	peak := 0
	for ; peak < len(arr)-1 && arr[peak+1] > arr[peak]; peak++ {
	}

	if peak == 0 || peak == len(arr)-1 {
		return false
	}

	for i := peak; i < len(arr)-1; i++ {
		if arr[i+1] >= arr[i] {
			return false
		}
	}

	return true
}
