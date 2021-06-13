package _336_palindrome_pairs

type Trie struct {
	Children           [26]*Trie
	Index              int
	FurtherPalindromes []int
}

func palindromePairs(words []string) [][]int {
	root := newTrieNode()
	for i, word := range words {
		node := insert(root, word, i)
		node.Index = i
	}

	result := [][]int{}
	for i, word := range words {
		cur := root
		if cur.Index != -1 && cur.Index != i {
			if isPalindrome(word) {
				result = append(result, []int{i, cur.Index})
			}
		}

		for j := 0; j < len(word); j++ {
			idx := word[j] - 'a'
			cur = cur.Children[idx]
			if cur == nil {
				break
			}

			if cur.Index != -1 && cur.Index != i && isPalindrome(word[len(words[cur.Index]):]) {
				result = append(result, []int{i, cur.Index})
			}
		}

		if cur != nil {
			for _, j := range cur.FurtherPalindromes {
				if i != j {
					result = append(result, []int{i, j})
				}
			}
		}
	}

	return result
}

func newTrieNode() *Trie {
	return &Trie{[26]*Trie{}, -1, nil}
}

func insert(node *Trie, word string, wordIdx int) *Trie {
	for i := len(word) - 1; i >= 0; i-- {
		if isPalindrome(word[:i+1]) {
			node.FurtherPalindromes = append(node.FurtherPalindromes, wordIdx)
		}

		idx := word[i] - 'a'
		if node.Children[idx] == nil {
			node.Children[idx] = newTrieNode()
		}
		node = node.Children[idx]
	}
	return node
}

func isPalindrome(word string) bool {
	for l, r := 0, len(word)-1; l < r; l, r = l+1, r-1 {
		if word[l] != word[r] {
			return false
		}
	}
	return true
}
