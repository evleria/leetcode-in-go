package _131_palindrome_partitioning

func partition(s string) [][]string {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	var res [][]string
	backtrack(s, 0, &res, []string{}, dp)
	return res
}

func backtrack(s string, start int, res *[][]string, current []string, dp [][]bool) {
	if start >= len(s) {
		currentCopy := make([]string, len(current))
		copy(currentCopy, current)
		*res = append(*res, currentCopy)
		return
	}
	for end := start; end < len(s); end++ {
		if s[start] == s[end] && (end-start <= 1 || dp[start+1][end-1]) {
			dp[start][end] = true
			current = append(current, s[start:end+1])
			backtrack(s, end+1, res, current, dp)
			current = current[:len(current)-1]
		}
	}
}
