package _752_open_the_lock

import (
	"fmt"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestOpenLock(t *testing.T) {
	testCases := []struct {
		gotD []string
		gotT string
		want int
	}{
		{
			gotD: []string{"0201", "0101", "0102", "1212", "2002"},
			gotT: "0202",
			want: 6,
		},
		{
			gotD: []string{"8888"},
			gotT: "0009",
			want: 1,
		},
		{
			gotD: []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"},
			gotT: "8888",
			want: -1,
		},
		{
			gotD: []string{"0000"},
			gotT: "8888",
			want: -1,
		},
		{
			gotD: []string{"1111"},
			gotT: "0000",
			want: 0,
		},
	}

	for _, testCase := range testCases {
		actual := openLock(testCase.gotD, testCase.gotT)

		assert.Check(t, is.DeepEqual(actual, testCase.want), fmt.Sprintf("%#v", testCase))
	}
}
