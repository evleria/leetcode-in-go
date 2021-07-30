package _677_map_sum_pairs

type Trie struct {
	Children map[int]*Trie
	Val      int
}

func NewTrie() *Trie {
	return &Trie{Children: map[int]*Trie{}}
}

type MapSum struct {
	Root *Trie
}

func Constructor() MapSum {
	return MapSum{
		Root: NewTrie(),
	}
}

func (m *MapSum) Insert(key string, val int) {
	m.navigateTo(key).Val = val
}

func (m *MapSum) Sum(prefix string) int {
	node := m.navigateTo(prefix)
	return sum(node)
}

func (m *MapSum) navigateTo(s string) *Trie {
	cur := m.Root
	for _, r := range s {
		idx := int(r - 'a')
		if _, ok := cur.Children[idx]; !ok {
			cur.Children[idx] = NewTrie()
		}
		cur = cur.Children[idx]
	}
	return cur
}

func sum(node *Trie) int {
	s := node.Val
	for _, n := range node.Children {
		if n != nil {
			s += sum(n)
		}
	}
	return s
}
