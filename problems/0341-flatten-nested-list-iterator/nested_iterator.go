package _341_flatten_nested_list_iterator

type NestedIterator struct {
	idx            int
	nestedIterator *NestedIterator
	nestedList     []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		idx:        -1,
		nestedList: nestedList,
	}
}

func (ni *NestedIterator) Next() int {
	if nestedIterator := ni.nestedIterator; nestedIterator != nil {
		return nestedIterator.Next()
	}

	return ni.nestedList[ni.idx].GetInteger()
}

func (ni *NestedIterator) HasNext() bool {
	if iterator := ni.nestedIterator; iterator != nil {
		if iterator.HasNext() {
			return true
		}
		ni.nestedIterator = nil
	}

	ni.idx++
	if ni.idx == len(ni.nestedList) {
		return false
	}

	if ni.nestedList[ni.idx].IsInteger() {
		return true
	}
	ni.nestedIterator = Constructor(ni.nestedList[ni.idx].GetList())
	return ni.HasNext()
}
