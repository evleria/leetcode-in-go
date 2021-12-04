package _032_stream_of_characters

type Trie struct {
	Children [26]*Trie
	IsEnd    bool
}

func (t *Trie) AddWord(word string) {
	cur := t
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if cur.Children[idx] == nil {
			cur.Children[idx] = new(Trie)
		}
		cur = cur.Children[idx]
	}
	cur.IsEnd = true
}

type StreamChecker struct {
	root  *Trie
	heads []*Trie
}

func Constructor(words []string) StreamChecker {
	root := new(Trie)
	for _, word := range words {
		root.AddWord(word)
	}

	return StreamChecker{
		root: root,
	}
}

func (sc *StreamChecker) Query(letter byte) bool {
	found := false
	sc.heads = append(sc.heads, sc.root)
	for _, head := range sc.heads {
		sc.heads = sc.heads[1:]
		idx := letter - 'a'
		if x := head.Children[idx]; x != nil {
			sc.heads = append(sc.heads, x)
			if x.IsEnd {
				found = true
			}
		}
	}

	return found
}
