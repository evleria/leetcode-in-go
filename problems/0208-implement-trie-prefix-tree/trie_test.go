package _208_implement_trie_prefix_tree

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	assert.Check(t, is.Equal(trie.Search("apple"), true))
	assert.Check(t, is.Equal(trie.Search("app"), false))
	assert.Check(t, is.Equal(trie.StartsWith("app"), true))
	trie.Insert("app")
	assert.Check(t, is.Equal(trie.Search("app"), true))
}
