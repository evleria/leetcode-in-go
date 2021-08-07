package _132_palindrome_partitioning_II

func minCut(s string) int {
	n := len(s)
	cut := make([]int, n+1)
	for i := 0; i <= n; i++ {
		cut[i] = i - 1
	}
	for i := 0; i < n; i++ {
		for j := 0; i-j >= 0 && i+j < n && s[i-j] == s[i+j]; j++ {
			cut[i+j+1] = min(cut[i+j+1], 1+cut[i-j])
		}
		for j := 1; i-j+1 >= 0 && i+j < n && s[i-j+1] == s[i+j]; j++ {
			cut[i+j+1] = min(cut[i+j+1], 1+cut[i-j+1])
		}
	}
	return cut[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
