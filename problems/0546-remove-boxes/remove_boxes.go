package _546_remove_boxes

func removeBoxes(boxes []int) int {
	n := len(boxes)
	dp := make([]int, n*n*n)
	return helper(boxes, 0, n-1, 0, n, dp)
}

func helper(boxes []int, i, j, k, n int, dp []int) int {
	if i > j {
		return 0
	}

	idx := k*n*n + j*n + i
	if r := dp[idx]; r > 0 {
		return r
	}

	max := (k+1)*(k+1) + helper(boxes, i+1, j, 0, n, dp)
	for l := i + 1; l <= j; l++ {
		if boxes[i] == boxes[l] {
			x := helper(boxes, i+1, l-1, 0, n, dp) + helper(boxes, l, j, k+1, n, dp)
			if x > max {
				max = x
			}
		}
	}
	dp[idx] = max
	return max
}
