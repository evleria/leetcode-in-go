package _926_flip_string_to_monotone_increasing

func minFlipsMonoIncr(s string) int {
	dp0, dp1 := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		dp0 = int(s[i]-'0') + min(dp0, dp1)
		dp1 = int('1'-s[i]) + dp1
	}
	return min(dp0, dp1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
