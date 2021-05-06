package _221_split_a_string_in_balanced_strings

func balancedStringSplit(s string) int {
	letter, count := 0, 0
	for i := range s {
		if s[i] == 'R' {
			letter++
		} else {
			letter--
		}
		if letter == 0 {
			count++
		}
	}
	return count
}
