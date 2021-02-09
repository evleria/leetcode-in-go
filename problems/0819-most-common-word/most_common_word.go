package _819_most_common_word

import (
	"strings"
	"unicode"
)

type Trie struct {
	Children [26]*Trie
	IsBanned bool
	Count    int
}

func mostCommonWord(paragraph string, banned []string) string {
	root := &Trie{}
	for _, word := range banned {
		node := insertWord(root, word)
		node.IsBanned = true
	}

	paragraph = strings.ToLower(paragraph)
	words := strings.FieldsFunc(paragraph, func(c rune) bool {
		return !unicode.IsLetter(c)
	})

	result, max := "", 0
	for _, word := range words {
		if node := insertWord(root, word); !node.IsBanned {
			node.Count++
			if node.Count > max {
				result, max = word, node.Count
			}
		}
	}

	return result
}

func insertWord(root *Trie, word string) *Trie {
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if root.Children[idx] == nil {
			root.Children[idx] = &Trie{}
		}

		root = root.Children[idx]
	}

	return root
}
