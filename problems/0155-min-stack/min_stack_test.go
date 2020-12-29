package _155_min_stack

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMinStack(t *testing.T) {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	assert.Check(t, is.Equal(stack.GetMin(), -3))
	stack.Pop()
	assert.Check(t, is.Equal(stack.Top(), 0))
	assert.Check(t, is.Equal(stack.GetMin(), -2))
}
