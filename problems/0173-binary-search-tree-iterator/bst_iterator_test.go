package _173_binary_search_tree_iterator

import (
	. "github.com/evleria/leetcode-in-go/structs/tree"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestBstIterator(t *testing.T) {
	tree := BinaryTreeFromSlice([]int{7, 3, 15, NULL, NULL, 9, 20})
	param := Constructor(tree)
	assert.Check(t, is.Equal(param.Next(), 3))
	assert.Check(t, is.Equal(param.Next(), 7))
	assert.Check(t, is.Equal(param.HasNext(), true))
	assert.Check(t, is.Equal(param.Next(), 9))
	assert.Check(t, is.Equal(param.HasNext(), true))
	assert.Check(t, is.Equal(param.Next(), 15))
	assert.Check(t, is.Equal(param.HasNext(), true))
	assert.Check(t, is.Equal(param.Next(), 20))
	assert.Check(t, is.Equal(param.HasNext(), false))
}
