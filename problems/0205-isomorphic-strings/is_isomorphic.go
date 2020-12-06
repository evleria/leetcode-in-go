package _205_isomorphic_strings

func isIsomorphic(s string, t string) bool {
	dict := make(map[byte]byte, len(s))
	seen := make(map[byte]bool, len(s))
	for i := 0; i < len(s); i++ {
		if x, ok := dict[s[i]]; ok {
			if t[i] != x {
				return false
			}
		} else {
			if _, ok := seen[t[i]]; ok {
				return false
			}
			seen[t[i]] = true
			dict[s[i]] = t[i]
		}
	}
	return true
}
