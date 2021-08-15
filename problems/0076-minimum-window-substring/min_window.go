package _076_minimum_window_substring

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	m := make(map[byte]int, 52)
	for i := 0; i < len(t); i++ {
		m[t[i]]--
	}
	n := len(m)

	result, start, end := "", 0, 0
	for ; n > 0 && end < len(s); end++ {
		ch := s[end]
		if x, ok := m[ch]; ok && x == -1 {
			n--
		}
		m[ch]++
		if n == 0 {
			break
		}
	}

	for end < len(s) {
		if l := end - start + 1; result == "" || l < len(result) {
			result = s[start : end+1]
		}

		if m[s[start]] > 0 {
			m[s[start]]--
			start++
		} else if end == len(s)-1 {
			break
		} else {
			m[s[end+1]]++
			end++
		}
	}

	return result
}
