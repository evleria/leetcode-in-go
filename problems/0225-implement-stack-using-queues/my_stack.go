package _225_implement_stack_using_queues

type MyStack struct {
	queue []int
	top   int
	size  int
}

func Constructor() MyStack {
	return MyStack{}
}

func (m *MyStack) Push(x int) {
	m.queue = append(m.queue, x)
	m.top = x
	m.size++
}

func (m *MyStack) Pop() int {
	s := m.size
	for i := 0; i < s-1; i++ {
		t := m.queue[0]
		m.queue = m.queue[1:]
		m.size--

		m.Push(t)
	}

	r := m.queue[0]
	m.queue = m.queue[1:]
	m.size--
	return r
}

func (m *MyStack) Top() int {
	return m.top
}

func (m *MyStack) Empty() bool {
	return m.size == 0
}
