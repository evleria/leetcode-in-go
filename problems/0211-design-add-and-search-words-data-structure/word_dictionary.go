package _211_design_add_and_search_words_data_structure

type trie struct {
	children [26]*trie
	isEnd    bool
}

func (t *trie) contains(word string) bool {
	if len(word) == 0 {
		return t.isEnd
	}

	first, rest := word[0], word[1:]
	if first == '.' {
		for _, child := range t.children {
			if child != nil {
				if child.contains(rest) {
					return true
				}
			}
		}
		return false
	}

	idx := first - 'a'
	if t.children[idx] == nil {
		return false
	}
	return t.children[idx].contains(rest)
}

type WordDictionary struct {
	root *trie
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &trie{},
	}
}

func (wd *WordDictionary) AddWord(word string) {
	cur := wd.root
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if cur.children[idx] == nil {
			cur.children[idx] = &trie{}
		}
		cur = cur.children[idx]
	}
	cur.isEnd = true
}

func (wd *WordDictionary) Search(word string) bool {
	return wd.root.contains(word)
}
