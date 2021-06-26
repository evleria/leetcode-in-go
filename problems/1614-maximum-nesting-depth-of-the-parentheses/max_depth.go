package _614_maximum_nesting_depth_of_the_parentheses

func maxDepth(s string) int {
	max, depth := 0, 0
	for _, ch := range s {
		if ch == '(' {
			depth++
			if depth > max {
				max = depth
			}
		} else if ch == ')' {
			depth--
		}
	}
	return max
}
