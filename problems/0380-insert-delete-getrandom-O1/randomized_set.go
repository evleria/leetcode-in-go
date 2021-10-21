package _380_insert_delete_getrandom_O1

import (
	"math/rand"
)

type RandomizedSet struct {
	m map[int]int
	n []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m: map[int]int{},
	}
}

func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.m[val]; !ok {
		r.m[val] = len(r.n)
		r.n = append(r.n, val)
		return true
	}
	return false
}

func (r *RandomizedSet) Remove(val int) bool {
	if i, ok := r.m[val]; ok {
		r.n[i] = r.n[len(r.n)-1]
		r.m[r.n[i]] = i
		delete(r.m, val)
		r.n = r.n[:len(r.n)-1]
		return true
	}
	return false
}

func (r *RandomizedSet) GetRandom() int {
	return r.n[rand.Intn(len(r.n))]
}
