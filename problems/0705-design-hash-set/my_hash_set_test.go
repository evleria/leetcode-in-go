package _705_design_hash_set

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMyHashSet(t *testing.T) {
	set := Constructor()
	set.Add(1)
	set.Add(2)
	assert.Check(t, is.Equal(set.Contains(1), true))
	assert.Check(t, is.Equal(set.Contains(3), false))
	set.Add(2)
	assert.Check(t, is.Equal(set.Contains(2), true))
	set.Remove(2)
	assert.Check(t, is.Equal(set.Contains(2), false))
}
