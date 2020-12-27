package _146_lru_cache

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestLruCache(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	assert.Check(t, is.Equal(cache.Get(1), 1))
	cache.Put(3, 3)
	assert.Check(t, is.Equal(cache.Get(2), -1))
	cache.Put(4, 4)
	assert.Check(t, is.Equal(cache.Get(1), -1))
	assert.Check(t, is.Equal(cache.Get(3), 3))
	assert.Check(t, is.Equal(cache.Get(4), 4))

	cache = Constructor(1)
	cache.Put(2, 1)
	assert.Check(t, is.Equal(cache.Get(2), 1))
	cache.Put(3, 2)
	assert.Check(t, is.Equal(cache.Get(2), -1))
	assert.Check(t, is.Equal(cache.Get(3), 2))
}
