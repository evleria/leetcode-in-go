package _456_132_pattern

type data struct {
	val int
	min int
}

func find132pattern(nums []int) bool {
	stack := make([]data, 0)
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		n := nums[i]
		for len(stack) > 0 && n >= stack[len(stack)-1].val {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 && n < stack[len(stack)-1].val && n > stack[len(stack)-1].min {
			return true
		}
		stack = append(stack, data{n, min})
		if n < min {
			min = n
		}
	}
	return false
}
