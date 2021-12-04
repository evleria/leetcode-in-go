package _032_stream_of_characters

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestStreamChecker(t *testing.T) {
	checker := Constructor([]string{"cd", "f", "kl"})
	assert.Check(t, is.Equal(checker.Query('a'), false))
	assert.Check(t, is.Equal(checker.Query('b'), false))
	assert.Check(t, is.Equal(checker.Query('c'), false))
	assert.Check(t, is.Equal(checker.Query('d'), true))
	assert.Check(t, is.Equal(checker.Query('e'), false))
	assert.Check(t, is.Equal(checker.Query('f'), true))
	assert.Check(t, is.Equal(checker.Query('g'), false))
	assert.Check(t, is.Equal(checker.Query('h'), false))
	assert.Check(t, is.Equal(checker.Query('i'), false))
	assert.Check(t, is.Equal(checker.Query('j'), false))
	assert.Check(t, is.Equal(checker.Query('k'), false))
	assert.Check(t, is.Equal(checker.Query('l'), true))
}
