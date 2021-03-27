package _647_palindromic_substrings

func countSubstrings(s string) int {
	count := 0
	for start := 0; start < len(s); start++ {
		count++
		left, right := start, start
		for ; right < len(s)-1 && s[right] == s[right+1]; right++ {
			count++
		}
		for ; left > 0 && right < len(s)-1 && s[left-1] == s[right+1]; left, right = left-1, right+1 {
			count++
		}
	}
	return count
}
