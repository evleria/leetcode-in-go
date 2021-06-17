package _387_first_unique_character_in_a_string

func firstUniqChar(s string) int {
	m := [26]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if m[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
