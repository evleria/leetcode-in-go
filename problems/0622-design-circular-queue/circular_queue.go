package _622_design_circular_queue

type MyCircularQueue struct {
	slice             []int
	front, rear, size int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		size:  k,
		slice: make([]int, k),
		front: 0,
		rear:  -1,
	}
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}
	q.rear++
	q.slice[q.rear%q.size] = value
	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	q.front++
	return true
}

func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.slice[q.front%q.size]
}

func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.slice[q.rear%q.size]
}

func (q *MyCircularQueue) IsEmpty() bool {
	return q.rear < q.front
}

func (q *MyCircularQueue) IsFull() bool {
	return q.rear-q.front == q.size-1
}
