package mergetwosortedlist

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_makeList(t *testing.T) {
	tests := [][]int{
		{},
		{1},
		{1, 2, 3, 4, 5},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			list := makeList(test)
			actual := listToSlice(list)
			require.Equal(t, test, actual)
			require.Len(t, actual, len(test))
		})
	}
}

func Test_copyList(t *testing.T) {
	tests := []struct {
		list   []int
		result []int
	}{
		{list: nil, result: []int{}},
		{list: []int{1}, result: []int{1}},
		{list: []int{1, 2, 3}, result: []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.list), func(t *testing.T) {
			list := makeList(tt.list)
			res := copyList(list)
			actual := listToSlice(res)
			require.Equal(t, tt.result, actual)
			require.Len(t, actual, len(tt.list))
		})
	}
}
