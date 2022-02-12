package _127_word_ladder

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if !contains(wordList, beginWord) {
		wordList = append(wordList, beginWord)
	}
	// construct our graph
	adjacencyList := make(map[string][]string, len(wordList))
	for i := 0; i < len(wordList)-1; i++ {
		for j := i + 1; j < len(wordList); j++ {
			if areTransformable(wordList[i], wordList[j]) {
				adjacencyList[wordList[i]] = append(adjacencyList[wordList[i]], wordList[j])
				adjacencyList[wordList[j]] = append(adjacencyList[wordList[j]], wordList[i])
			}
		}
	}

	visited, queue, pathLen := make(map[string]bool, len(wordList)), []string{beginWord}, 1
	for len(queue) > 0 {
		pathLen++
		for _, current := range queue {
			queue = queue[1:]
			visited[current] = true

			for _, candidate := range adjacencyList[current] {
				if !visited[candidate] {
					if candidate == endWord {
						return pathLen
					}

					queue = append(queue, candidate)
				}
			}
		}
	}
	return 0
}

func areTransformable(a, b string) bool {
	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff == 1
}

func contains(in []string, s string) bool {
	for _, w := range in {
		if w == s {
			return true
		}
	}
	return false
}
