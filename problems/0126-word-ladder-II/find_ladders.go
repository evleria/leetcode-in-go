package _126_word_ladder_II

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
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

	queue, result := [][]string{{beginWord}}, [][]string{}
	pathLen := make(map[string]int, len(wordList))
	pathLen[beginWord] = 1

	// BFS processing
	for len(queue) > 0 && len(result) == 0 {
		for _, path := range queue {
			queue = queue[1:]
			// if we already know any shorter path to 'candidate' word - drop current path
			for _, candidate := range adjacencyList[path[len(path)-1]] {
				if l, ok := pathLen[candidate]; ok && l < len(path)+1 {
					continue
				}

				newPath := append(append(make([]string, 0, len(path)+1), path...), candidate)
				if candidate == endWord {
					result = append(result, newPath)
				} else {
					queue = append(queue, newPath)
					pathLen[candidate] = len(newPath)
				}
			}
		}
	}

	return result
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
