package _978_longest_turbulent_subarray

func maxTurbulenceSize(arr []int) int {
	left, prevSign, result := 0, 0, 1
	for right := 1; right < len(arr); right++ {
		sign := getSign(arr[right] - arr[right-1])
		switch {
		case sign == 0:
			left = right
		case sign == prevSign:
			left = right - 1
		default:
			if d := right - left + 1; d > result {
				result = d
			}
		}

		prevSign = sign
	}

	return result
}

func getSign(a int) int {
	if a == 0 {
		return a
	} else if a < 0 {
		return -1
	}
	return 1
}
