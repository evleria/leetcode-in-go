package _209_remove_all_adjacent_duplicates_in_string_II

type Frame struct {
	Ch    byte
	Count int
}

func removeDuplicates(s string, k int) string {
	var stack []Frame
	for i := 0; i < len(s); i++ {
		top := len(stack) - 1
		if top >= 0 && stack[top].Ch == s[i] {
			stack[top].Count++
			if stack[top].Count == k {
				stack = stack[:top]
			}
		} else {
			stack = append(stack, Frame{s[i], 1})
		}
	}

	var result []byte
	for i := 0; i < len(stack); i++ {
		ch, count := stack[i].Ch, stack[i].Count
		for j := 0; j < count; j++ {
			result = append(result, ch)
		}
	}

	return string(result)
}
