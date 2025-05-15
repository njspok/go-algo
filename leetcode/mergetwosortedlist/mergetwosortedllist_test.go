package mergetwosortedlist

import (
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		a      []int
		b      []int
		result []int
	}{
		{a: nil, b: nil, result: []int{}},
		{a: []int{1, 2, 3}, b: nil, result: []int{1, 2, 3}},
		{a: nil, b: []int{4, 5, 6}, result: []int{4, 5, 6}},
		{a: []int{1, 2, 3}, b: []int{0}, result: []int{0, 1, 2, 3}},
		{a: []int{1, 2, 3}, b: []int{1}, result: []int{1, 1, 2, 3}},
		{a: []int{1, 2, 3}, b: []int{2}, result: []int{1, 2, 2, 3}},
		{a: []int{1, 2, 3}, b: []int{3}, result: []int{1, 2, 3, 3}},
		{a: []int{1, 2, 3}, b: []int{4}, result: []int{1, 2, 3, 4}},
		{a: []int{1, 2, 3}, b: []int{4, 5}, result: []int{1, 2, 3, 4, 5}},
		{a: []int{1, 2, 3}, b: []int{4, 5, 6}, result: []int{1, 2, 3, 4, 5, 6}},
		{a: []int{4, 5, 6}, b: []int{1, 2, 3}, result: []int{1, 2, 3, 4, 5, 6}},
		{a: []int{1, 2, 4}, b: []int{1, 3, 4}, result: []int{1, 1, 2, 3, 4, 4}},
	}
	for n, tt := range tests {
		t.Run(strconv.Itoa(n), func(t *testing.T) {
			res := mergeTwoLists(makeList(tt.a), makeList(tt.b))
			actual := listToSlice(res)
			require.Equal(t, tt.result, actual)
		})
	}
}
