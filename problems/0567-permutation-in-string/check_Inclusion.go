package _567_permutation_in_string

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	diff := [26]int{}
	for i := 0; i < len(s1); i++ {
		diff[s1[i]-'a']++
		diff[s2[i]-'a']--
	}

	empty := [26]int{}
	for i := 0; i < len(s2)-len(s1); i++ {
		if diff == empty {
			return true
		}
		diff[s2[i]-'a']++
		diff[s2[i+len(s1)]-'a']--
	}

	return diff == empty
}
