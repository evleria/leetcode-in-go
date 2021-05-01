package _745_prefix_and_suffix_search

import (
	"fmt"
)

type TrieNode struct {
	Index    int
	Children *[27]*TrieNode
}

type WordFilter struct {
	root *TrieNode
}

func Constructor(words []string) WordFilter {
	root := &TrieNode{}
	for i, w := range words {
		for j := 0; j < len(w); j++ {
			root.insert(w[j:], i).insert("{", i).insert(w, i)
		}
	}

	return WordFilter{
		root: root,
	}
}

func (wf *WordFilter) F(prefix string, suffix string) int {
	n := wf.root.find(fmt.Sprintf("%s{%s", suffix, prefix))
	if n == nil {
		return -1
	}
	return n.Index
}

func (tn *TrieNode) insert(s string, wordIndex int) *TrieNode {
	current := tn
	for i := 0; i < len(s); i++ {
		if current.Children == nil {
			current.Children = &[27]*TrieNode{}
		}

		idx := s[i] - 'a'
		if current.Children[idx] == nil {
			current.Children[idx] = &TrieNode{}
		}
		current = current.Children[idx]
		current.Index = wordIndex
	}
	return current
}

func (tn *TrieNode) find(s string) *TrieNode {
	current := tn
	for i := 0; i < len(s); i++ {
		if current.Children == nil {
			return nil
		}
		idx := s[i] - 'a'
		if current.Children[idx] == nil {
			return nil
		}
		current = current.Children[idx]
	}
	return current
}
