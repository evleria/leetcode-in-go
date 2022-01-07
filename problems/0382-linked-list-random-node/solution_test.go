package _382_linked_list_random_node

import (
	. "github.com/evleria/leetcode-in-go/structs/linked-list"
	"math/rand"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestSolution(t *testing.T) {
	rand.Seed(100)

	list := LinkedListFromSlice([]int{1, 2, 3})
	head := Constructor(list)
	assert.Check(t, is.Equal(head.GetRandom(), 2))
	assert.Check(t, is.Equal(head.GetRandom(), 3))
	assert.Check(t, is.Equal(head.GetRandom(), 2))
	assert.Check(t, is.Equal(head.GetRandom(), 1))
	assert.Check(t, is.Equal(head.GetRandom(), 1))
}
