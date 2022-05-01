package _844_backspace_string_compare

func backspaceCompare(s string, t string) bool {
	if compare(s) == compare(t) {
		return true
	} else {
		return false
	}
}

func compare(s string) string {
	str := []byte{}
	for _, c := range s {
		if c == '#' {
			if len(str) > 0 {
				str = str[:len(str)-1]
			}
		} else {
			str = append(str, byte(c))
		}
	}
	return string(str[:])
}
