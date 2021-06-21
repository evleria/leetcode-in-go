package _528_shuffle_string

func restoreString(s string, indices []int) string {
	result := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		result[indices[i]] = s[i]
	}
	return string(result)
}
