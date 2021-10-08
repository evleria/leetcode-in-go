package _208_implement_trie_prefix_tree

type Trie struct {
	Children [26]*Trie
	IsEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	if len(word) > 0 {
		idx := word[0] - 'a'
		if t.Children[idx] == nil {
			t.Children[idx] = &Trie{}
		}
		t.Children[idx].Insert(word[1:])
	} else {
		t.IsEnd = true
	}
}

func (t *Trie) Search(word string) bool {
	if len(word) > 0 {
		idx := word[0] - 'a'
		if t.Children[idx] == nil {
			return false
		} else {
			return t.Children[idx].Search(word[1:])
		}
	}
	return t.IsEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	if len(prefix) > 0 {
		idx := prefix[0] - 'a'
		if t.Children[idx] == nil {
			return false
		} else {
			return t.Children[idx].StartsWith(prefix[1:])
		}
	}
	return true
}
