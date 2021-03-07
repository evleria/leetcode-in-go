package _706_design_hashmap

import (
	"math/rand"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestMyHashMap(t *testing.T) {
	hashMap := Constructor()
	hashMap.Put(1, 1)
	hashMap.Put(2, 2)
	assert.Check(t, is.Equal(hashMap.Get(1), 1))
	assert.Check(t, is.Equal(hashMap.Get(3), -1))
	hashMap.Put(2, 1)
	assert.Check(t, is.Equal(hashMap.Get(2), 1))
	hashMap.Remove(2)
	assert.Check(t, is.Equal(hashMap.Get(2), -1))
}

func TestResizing(t *testing.T) {
	builtinMap := map[int]int{}
	hashMap := Constructor()

	for i := 0; i < 1000; i++ {
		k, v := rand.Intn(1000), rand.Intn(1000)
		hashMap.Put(k, v)
		builtinMap[k] = v
	}

	for k := 0; k < 1000; k++ {
		if v, ok := builtinMap[k]; ok {
			assert.Check(t, is.Equal(hashMap.Get(k), v), k)
		} else {
			assert.Check(t, is.Equal(hashMap.Get(k), -1), k)
		}
	}
}
