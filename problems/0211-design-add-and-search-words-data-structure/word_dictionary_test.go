package _211_design_add_and_search_words_data_structure

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestWordDictionary(t *testing.T) {
	dictionary := Constructor()
	dictionary.AddWord("bad")
	dictionary.AddWord("dad")
	dictionary.AddWord("mad")
	assert.Check(t, is.Equal(dictionary.Search("pad"), false))
	assert.Check(t, is.Equal(dictionary.Search("bad"), true))
	assert.Check(t, is.Equal(dictionary.Search(".ad"), true))
	assert.Check(t, is.Equal(dictionary.Search("b.."), true))
}
