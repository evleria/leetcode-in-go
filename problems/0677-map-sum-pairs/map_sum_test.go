package _677_map_sum_pairs

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMapSum(t *testing.T) {
	mapSum := Constructor()
	mapSum.Insert("apple", 3)
	assert.Check(t, is.Equal(mapSum.Sum("ap"), 3))
	mapSum.Insert("app", 2)
	assert.Check(t, is.Equal(mapSum.Sum("ap"), 5))
}
