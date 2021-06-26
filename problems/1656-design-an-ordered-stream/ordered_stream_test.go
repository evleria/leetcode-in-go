package _656_design_an_ordered_stream

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestOrderedStream(t *testing.T) {
	orderedStream := Constructor(5)
	assert.Check(t, is.DeepEqual(orderedStream.Insert(3, "ccccc"), []string{}))
	assert.Check(t, is.DeepEqual(orderedStream.Insert(1, "aaaaa"), []string{"aaaaa"}))
	assert.Check(t, is.DeepEqual(orderedStream.Insert(2, "bbbbb"), []string{"bbbbb", "ccccc"}))
	assert.Check(t, is.DeepEqual(orderedStream.Insert(5, "eeeee"), []string{}))
	assert.Check(t, is.DeepEqual(orderedStream.Insert(4, "ddddd"), []string{"ddddd", "eeeee"}))
}
