package _895_maximum_frequency_stack

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestFreqStack(t *testing.T) {
	freqStack := Constructor()
	freqStack.Push(5)
	freqStack.Push(7)
	freqStack.Push(5)
	freqStack.Push(7)
	freqStack.Push(4)
	freqStack.Push(5)
	assert.Check(t, is.Equal(freqStack.Pop(), 5))
	assert.Check(t, is.Equal(freqStack.Pop(), 7))
	assert.Check(t, is.Equal(freqStack.Pop(), 5))
	assert.Check(t, is.Equal(freqStack.Pop(), 4))
}
