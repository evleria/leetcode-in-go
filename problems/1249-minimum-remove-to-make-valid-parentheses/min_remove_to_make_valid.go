package _249_minimum_remove_to_make_valid_parentheses

func minRemoveToMakeValid(s string) string {
	open, r := []int{}, make([]rune, 0, len(s))

	for _, ch := range s {
		switch ch {
		case '(':
			open = append(open, len(r))
		case ')':
			if len(open) > 0 {
				open = open[:len(open)-1]
			} else {
				continue
			}
		}
		r = append(r, ch)
	}

	for i := len(open) - 1; i >= 0; i-- {
		idx := open[i]
		r = append(r[:idx], r[idx+1:]...)
	}
	return string(r)
}
