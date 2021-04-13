package _341_flatten_nested_list_iterator

import (
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestNestedIterator(t *testing.T) {
	testCases := []struct {
		got  *NestedInteger
		want []int
	}{
		{
			//[[]]
			got:  new(NestedInteger).Append(new(NestedInteger)),
			want: []int{},
		},
		{
			//[[],[3]]
			got: new(NestedInteger).Append(
				new(NestedInteger),
				new(NestedInteger).Append(Int(3))),
			want: []int{3},
		},
		{
			//[[1,1],2,[1,1]]
			got: new(NestedInteger).Append(
				new(NestedInteger).Append(Int(1), Int(1)),
				Int(2),
				new(NestedInteger).Append(Int(1), Int(1)),
			),
			want: []int{1, 1, 2, 1, 1},
		},
		{
			//[1,[4,[6]]]
			got: new(NestedInteger).Append(
				Int(1),
				new(NestedInteger).Append(
					Int(4),
					new(NestedInteger).Append(Int(6)))),
			want: []int{1, 4, 6},
		},
	}

	for _, testCase := range testCases {
		iterator := Constructor(testCase.got.GetList())
		list := []int{}
		for iterator.HasNext() {
			n := iterator.Next()
			list = append(list, n)
		}

		assert.Check(t, is.DeepEqual(list, testCase.want))
	}
}

func (ni *NestedInteger) Append(elem ...*NestedInteger) *NestedInteger {
	for _, i := range elem {
		ni.Add(*i)
	}
	return ni
}

func Int(v int) *NestedInteger {
	ni := new(NestedInteger)
	ni.SetInteger(v)
	return ni
}
