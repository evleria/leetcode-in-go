package _049_group_anagrams

func groupAnagrams(strs []string) [][]string {
	dict := map[[26]byte][]string{}
	for _, s := range strs {
		letters := [26]byte{}
		for i := 0; i < len(s); i++ {
			letters[s[i]-'a']++
		}
		dict[letters] = append(dict[letters], s)
	}

	var result [][]string
	for _, strings := range dict {
		result = append(result, strings)
	}

	return result
}
