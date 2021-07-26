package _920_build_array_from_permutation

func buildArray(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		nums[i] = nums[i] + (n * (nums[nums[i]] % n))
	}
	for i := 0; i < n; i++ {
		nums[i] /= n
	}
	return nums
}
