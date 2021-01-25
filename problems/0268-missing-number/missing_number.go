package _268_missing_number

func missingNumber(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	for _, n := range nums {
		sum -= n
	}
	return sum
}
