package _303_range_sum_query_immutable

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return NumArray{nums: nums}
}

func (n *NumArray) SumRange(left int, right int) int {
	sum := n.nums[right]
	if left > 0 {
		sum -= n.nums[left-1]
	}
	return sum
}
