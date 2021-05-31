package _268_search_suggestions_system

type Trie struct {
	Children [26]*Trie
	IsEnd    bool
}

func suggestedProducts(products []string, searchWord string) [][]string {
	root := &Trie{}
	for _, product := range products {
		insertWord(root, product).IsEnd = true
	}

	result := make([][]string, 0, len(searchWord))
	for i := 0; i < len(searchWord); i++ {
		if root != nil {
			root = root.Children[searchWord[i]-'a']
		}
		result = append(result, scanWords(root, searchWord[:i+1]))
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

func scanWords(root *Trie, prefix string) []string {
	if root == nil {
		return []string{}
	}

	result := make([]string, 0, 3)
	dfs(root, []byte(prefix), &result)
	return result
}

func dfs(node *Trie, path []byte, result *[]string) {
	if len(*result) == 3 {
		return
	}
	if node.IsEnd {
		*result = append(*result, string(path))
	}
	for ch := 'a'; ch <= 'z'; ch++ {
		if n := node.Children[ch-'a']; n != nil {
			dfs(n, append(path, byte(ch)), result)
		}
	}
}
