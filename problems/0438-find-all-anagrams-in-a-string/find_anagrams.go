package _438_find_all_anagrams_in_a_string

var empty = [26]int{}

func findAnagrams(s string, p string) []int {
	result := []int{}
	if len(p) > len(s) {
		return result
	}
	diff := [26]int{}
	for i := 0; i < len(p); i++ {
		diff[p[i]-'a']++
		diff[s[i]-'a']--
	}
	if diff == empty {
		result = append(result, 0)
	}
	for i := 1; i < len(s)-len(p)+1; i++ {
		diff[s[i-1]-'a']++
		diff[s[i+len(p)-1]-'a']--
		if diff == empty {
			result = append(result, i)
		}
	}

	return result
}
