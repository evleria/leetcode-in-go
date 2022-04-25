package _284_peeking_iterator

type PeekingIterator struct {
	iter Iterator
	n    *int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter: *iter}
}

func (pi *PeekingIterator) hasNext() bool {
	return pi.n != nil || pi.iter.hasNext()
}

func (pi *PeekingIterator) next() int {
	if pi.n != nil {
		n := *pi.n
		pi.n = nil
		return n
	}
	return pi.iter.next()
}

func (pi *PeekingIterator) peek() int {
	if pi.n == nil {
		n := pi.iter.next()
		pi.n = &n
	}
	return *pi.n
}
