package _341_flatten_nested_list_iterator

type NestedInteger struct {
	data interface{}
}

func (ni NestedInteger) IsInteger() bool {
	_, ok := ni.data.(int)
	return ok
}

func (ni NestedInteger) GetInteger() int {
	return ni.data.(int)
}

func (ni *NestedInteger) SetInteger(value int) {
	ni.data = value
}

func (ni *NestedInteger) Add(elem NestedInteger) {
	if ni.data == nil {
		ni.data = []*NestedInteger{&elem}
		return
	}

	ni.data = append(ni.data.([]*NestedInteger), &elem)
}

func (ni NestedInteger) GetList() []*NestedInteger {
	if ni.data == nil {
		return nil
	}

	return ni.data.([]*NestedInteger)
}
