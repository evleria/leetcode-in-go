package _330_patching_array

func minPatches(nums []int, n int) int {
	miss, added, i := 1, 0, 0
	for miss <= n {
		if i < len(nums) && nums[i] <= miss {
			miss += nums[i]
			i++
		} else {
			miss *= 2
			added++
		}
	}
	return added
}
