package _461_check_if_a_string_contains_all_binary_codes_of_size_k

func hasAllCodes(s string, k int) bool {
	max := 1 << k
	if max > len(s)-k+1 {
		return false
	}
	m := make(map[string]bool, max)
	for i := 0; i <= len(s)-k; i++ {
		m[s[i:i+k]] = true
		if len(m) == max {
			return true
		}
	}
	return false
}
