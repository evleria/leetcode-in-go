package _560_subarray_sum_equals_k

func subarraySum(nums []int, k int) int {
	count, sum, prefix := 0, 0, map[int]int{0: 1}
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		count += prefix[sum-k]
		prefix[sum]++
	}
	return count
}
