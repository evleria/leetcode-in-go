package _014_longest_common_prefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	first := strs[0]
	for i := 0; i < len(first); i++ {
		for _, str := range strs[1:] {
			if i >= len(str) || str[i] != first[i] {
				return first[:i]
			}
		}
	}

	return first
}
