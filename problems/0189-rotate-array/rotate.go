package _189_rotate_array

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n

	if k == 0 {
		return
	}

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start, end int) {
	for l, r := start, end; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}
