package _146_lru_cache

type LRUCache struct {
	head     *Node
	tail     *Node
	m        map[int]*Node
	capacity int
}

type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		head:     new(Node),
		tail:     new(Node),
		m:        make(map[int]*Node, capacity),
		capacity: capacity,
	}

	cache.head.Prev, cache.tail.Next = cache.tail, cache.head

	return cache
}

func (this *LRUCache) Get(key int) int {
	if n, ok := this.m[key]; ok {
		this.deattach(n)
		this.insertFirst(n)
		return n.Val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if n, ok := this.m[key]; ok {
		n.Val = value
		this.deattach(n)
		this.insertFirst(n)
		return
	}

	node := &Node{Key: key, Val: value}
	this.m[key] = node
	this.insertFirst(node)

	if len(this.m) > this.capacity {
		last := this.tail.Next
		delete(this.m, last.Key)
		this.deattach(last)
	}
}

func (this *LRUCache) insertFirst(node *Node) {
	this.head.Prev, this.head.Prev.Next, node.Next, node.Prev = node, node, this.head, this.head.Prev
}

func (this *LRUCache) deattach(n *Node) {
	n.Next.Prev, n.Prev.Next = n.Prev, n.Next
}
