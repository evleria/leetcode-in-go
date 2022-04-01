package _002_find_common_characters

func commonChars(words []string) []string {
	firstMap := make(map[rune]int)
	for _, ch := range words[0] {
		firstMap[ch]++
	}

	for i := 1; i < len(words); i++ {
		secondMap := make(map[rune]int)
		for _, ch := range words[i] {
			secondMap[ch]++
		}

		for k, v := range firstMap {
			if _, ok := secondMap[k]; ok {
				firstMap[k] = min(v, secondMap[k])
			} else {
				delete(firstMap, k)
			}
		}
	}

	result := []string{}
	for k, v := range firstMap {
		for i := 0; i < v; i++ {
			result = append(result, string(k))
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
