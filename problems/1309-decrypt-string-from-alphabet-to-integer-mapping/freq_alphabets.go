package _309_decrypt_string_from_alphabet_to_integer_mapping

import "strings"

func freqAlphabets(s string) string {
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		n := s[i] - '0'
		if i+2 < len(s) && s[i+2] == '#' {
			n = 10*n + s[i+1] - '0'
			i += 2
		}
		sb.WriteByte(n + 'a' - 1)
	}
	return sb.String()
}
