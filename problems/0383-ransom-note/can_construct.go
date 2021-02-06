package _383_ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	letters := [26]int{}
	for i := 0; i < len(magazine); i++ {
		letters[magazine[i]-'a']++
	}
	for i := 0; i < len(ransomNote); i++ {
		l := ransomNote[i] - 'a'
		if letters[l] == 0 {
			return false
		}
		letters[l]--
	}

	return true
}
