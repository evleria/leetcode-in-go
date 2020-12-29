package _155_min_stack

type MinStack struct {
	min   []int
	stack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(x int) {
	s.stack = append(s.stack, x)
	if len(s.min) == 0 || s.min[len(s.min)-1] >= x {
		s.min = append(s.min, x)
	}
}

func (s *MinStack) Pop() {
	n := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	if s.min[len(s.min)-1] == n {
		s.min = s.min[:len(s.min)-1]
	}
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.min[len(s.min)-1]
}
