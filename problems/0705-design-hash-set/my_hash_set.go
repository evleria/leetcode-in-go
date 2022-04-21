package _705_design_hash_set

const capacity = 1_000

type MyHashSet struct {
	buckets [capacity]*Node
}

type Node struct {
	key  int
	next *Node
}

func Constructor() MyHashSet {
	return MyHashSet{}
}

func (m *MyHashSet) Add(key int) {
	idx := key % capacity
	if m.buckets[idx] == nil {
		m.buckets[idx] = &Node{key, nil}
	} else {
		cur := m.buckets[idx]
		for cur.next != nil {
			if cur.key == key {
				return
			}
			cur = cur.next
		}
		if cur.key == key {
			return
		}
		cur.next = &Node{key, nil}
	}
}

func (m *MyHashSet) Remove(key int) {
	idx := key % capacity
	if m.buckets[idx] == nil {
		return
	}
	cur := m.buckets[idx]
	if cur.key == key {
		m.buckets[idx] = cur.next
		return
	}
	for cur.next != nil {
		if cur.next.key == key {
			cur.next = cur.next.next
			return
		}
		cur = cur.next
	}
}

func (m *MyHashSet) Contains(key int) bool {
	idx := key % capacity
	cur := m.buckets[idx]
	for cur != nil {
		if cur.key == key {
			return true
		}
		cur = cur.next
	}
	return false
}
