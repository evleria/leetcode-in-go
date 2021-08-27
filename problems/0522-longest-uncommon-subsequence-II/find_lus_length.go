package _522_longest_uncommon_subsequence_II

func findLUSlength(strs []string) int {
	res := -1
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < res {
			continue
		}
		j := 0
		for ; j < len(strs); j++ {
			if i != j && isSubsequence(strs[i], strs[j]) {
				break
			}
		}
		if j == len(strs) && len(strs[i]) > res {
			res = len(strs[i])
		}
	}
	return res
}

func isSubsequence(a, b string) bool {
	i := 0
	for j := 0; i < len(a) && j < len(b); j++ {
		if a[i] == b[j] {
			i++
		}
	}
	return i == len(a)
}
