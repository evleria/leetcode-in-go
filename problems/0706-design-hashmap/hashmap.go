package _706_design_hashmap

type kvp struct {
	key   int
	value int
}

type MyHashMap struct {
	buckets [][]kvp
	count   int
}

const Threshold = 1.5

func Constructor() MyHashMap {
	return MyHashMap{
		buckets: make([][]kvp, 16),
	}
}

func (m *MyHashMap) Put(key int, value int) {
	n := key % len(m.buckets)
	bucket := m.buckets[n]
	for i := 0; i < len(bucket); i++ {
		if bucket[i].key == key {
			bucket[i].value = value
			return
		}
	}
	m.buckets[n] = append(bucket, kvp{key: key, value: value})
	m.count++
	m.resizeIfNeeded()
}

func (m *MyHashMap) Get(key int) int {
	n := key % len(m.buckets)
	bucket := m.buckets[n]
	for _, pair := range bucket {
		if pair.key == key {
			return pair.value
		}
	}
	return -1
}

func (m *MyHashMap) Remove(key int) {
	n := key % len(m.buckets)
	bucket := m.buckets[n]
	for i, pair := range bucket {
		if pair.key == key {
			m.buckets[n] = append(bucket[:i], bucket[i+1:]...)
			m.count--
			return
		}
	}
}

func (m *MyHashMap) resizeIfNeeded() {
	if float64(m.count)/float64(len(m.buckets)) > Threshold {
		oldLen := len(m.buckets)
		m.buckets = append(m.buckets, make([][]kvp, len(m.buckets))...)
		for i := 0; i < oldLen; i++ {
			for j := len(m.buckets[i]) - 1; j >= 0; j-- {
				pair := m.buckets[i][j]
				if newI := pair.key % len(m.buckets); newI != i {
					m.buckets[newI] = append(m.buckets[newI], pair)
					m.buckets[i] = append(m.buckets[i][:j], m.buckets[i][j+1:]...)
				}
			}
		}
	}
}
