package _695_maximum_erasure_value

func maximumUniqueSubarray(nums []int) int {
	l, sum, max := 0, 0, 0
	m := map[int]int{}
	for r := 0; r < len(nums); r++ {
		if j, ok := m[nums[r]]; ok && j >= l {
			for ; l <= j; l++ {
				sum -= nums[l]
			}
		}
		sum += nums[r]
		m[nums[r]] = r
		if sum > max {
			max = sum
		}
	}
	return max
}
