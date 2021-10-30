package _044_longest_duplicate_substring

func longestDupSubstring(s string) string {
	longest := ""
	left, right := 1, len(s)-1
	for left <= right {
		mid := (left + right) / 2
		sub, ok := findDuplicate(s, mid)
		if ok {
			left, longest = mid+1, sub
		} else {
			right = mid - 1
		}
	}
	return longest
}

func findDuplicate(s string, length int) (string, bool) {
	seen := make(map[string]bool, len(s)-length)
	for i := 0; i < len(s)-length+1; i++ {
		sub := s[i : i+length]
		if _, ok := seen[sub]; ok {
			return sub, true
		} else {
			seen[sub] = true
		}
	}
	return "", false
}
