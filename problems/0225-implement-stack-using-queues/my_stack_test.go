package _225_implement_stack_using_queues

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMyStack(t *testing.T) {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	assert.Check(t, is.Equal(stack.Top(), 2))
	assert.Check(t, is.Equal(stack.Pop(), 2))
	assert.Check(t, is.Equal(stack.Empty(), false))
}
