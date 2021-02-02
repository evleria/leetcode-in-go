package _392_is_subsequence

func isSubsequence(s string, t string) bool {
	si := 0
	for i := 0; i < len(t) && si != len(s); i++ {
		if t[i] == s[si] {
			si++
		}
	}
	return si == len(s)
}
