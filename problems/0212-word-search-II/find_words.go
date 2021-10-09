package _212_word_search_II

var directions = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func findWords(board [][]byte, words []string) []string {
	trie := new(TrieNode)
	for _, word := range words {
		trie.insert(word)
	}

	result := make([]string, 0, len(words))
	for i, row := range board {
		for j, ch := range row {
			idx := ch - 'a'
			if tn := trie.Children[idx]; tn != nil {
				dfs(board, i, j, trie, &result)
			}
		}
	}

	return result
}

func dfs(board [][]byte, i, j int, node *TrieNode, result *[]string) {
	letter := board[i][j]
	node = node.Children[letter-'a']
	if node == nil {
		return
	} else if node.Word != "" {
		*result = append(*result, node.Word)
		node.Word = ""
	}

	board[i][j] = '.'
	for _, dir := range directions {
		x, y := i+dir[0], j+dir[1]
		if x >= 0 && x < len(board) && y >= 0 && y < len(board[0]) && board[x][y] != '.' {
			dfs(board, x, y, node, result)
		}
	}
	board[i][j] = letter
}

type TrieNode struct {
	Children [26]*TrieNode
	Word     string
}

func (tn *TrieNode) insert(word string) {
	for _, ch := range word {
		idx := ch - 'a'
		if tn.Children[idx] == nil {
			tn.Children[idx] = new(TrieNode)
		}
		tn = tn.Children[idx]
	}
	tn.Word = word
}
