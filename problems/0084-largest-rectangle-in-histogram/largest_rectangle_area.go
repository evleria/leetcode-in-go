package _084_largest_rectangle_in_histogram

func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := []int{}
	left := make([]int, n)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[i] <= heights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = []int{}
	right := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[i] <= heights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	res := 0
	area := 0
	for i := 0; i < n; i++ {
		h, l, r := heights[i], left[i], right[i]
		area = h * (r - l - 1)
		if area > res {
			res = area
		}
	}
	return res
}
