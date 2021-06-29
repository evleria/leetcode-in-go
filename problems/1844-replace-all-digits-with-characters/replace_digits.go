package _844_replace_all_digits_with_characters

func replaceDigits(s string) string {
	result := make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			result[i] = s[i]
		} else {
			result[i] = s[i-1] + s[i] - '0'
		}
	}
	return string(result)
}
