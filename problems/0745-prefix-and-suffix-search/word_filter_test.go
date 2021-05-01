package _745_prefix_and_suffix_search

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestWordFilter(t *testing.T) {
	filter := Constructor([]string{"apple", "pine", "ape"})
	assert.Check(t, is.Equal(filter.F("a", "e"), 2))
	assert.Check(t, is.Equal(filter.F("p", "e"), 1))
	assert.Check(t, is.Equal(filter.F("p", "a"), -1))
}
