package _916_word_subsets

func wordSubsets(A []string, B []string) []string {
	b := [26]byte{}
	for _, s := range B {
		ch := getChars(s)
		merge(&b, ch)
	}

	result := []string{}
	for _, s := range A {
		ch := getChars(s)
		if contains(ch, b) {
			result = append(result, s)
		}
	}
	return result
}

func getChars(s string) [26]byte {
	r := [26]byte{}
	for i := 0; i < len(s); i++ {
		r[s[i]-'a']++
	}
	return r
}

func merge(a *[26]byte, b [26]byte) {
	for i := 0; i < 26; i++ {
		if b[i] > a[i] {
			a[i] = b[i]
		}
	}
}

func contains(a, b [26]byte) bool {
	for i := 0; i < 26; i++ {
		if a[i] < b[i] {
			return false
		}
	}
	return true
}
