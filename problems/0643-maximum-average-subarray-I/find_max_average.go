package _643_maximum_average_subarray_I

func findMaxAverage(nums []int, k int) float64 {
	window := 0
	for i := 0; i < k; i++ {
		window += nums[i]
	}
	max := window
	for i := k; i < len(nums); i++ {
		window += nums[i] - nums[i-k]
		if window > max {
			max = window
		}
	}
	return float64(max) / float64(k)
}
