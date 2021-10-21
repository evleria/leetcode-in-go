package _380_insert_delete_getrandom_O1

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestRandomizedSet(t *testing.T) {
	randSet := Constructor()
	assert.Check(t, is.Equal(randSet.Insert(1), true))
	assert.Check(t, is.Equal(randSet.Remove(2), false))
	assert.Check(t, is.Equal(randSet.Insert(2), true))
	randSet.GetRandom()
	assert.Check(t, is.Equal(randSet.Remove(1), true))
	assert.Check(t, is.Equal(randSet.Insert(2), false))
	randSet.GetRandom()
}
