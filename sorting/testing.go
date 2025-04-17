package sorting

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"slices"
	"sort"
	"testing"
)

func TestSort(t *testing.T, sortFn func([]int)) {
	tests := []struct {
		list []int
	}{
		{list: nil},
		{list: []int{}},
		{list: []int{11}},
		{list: []int{11, 22}},
		{list: []int{22, 11}},
		{list: []int{33, 22, 11}},
		{list: []int{33, 22, 11, 44, 66, 33, 11}},
		{list: []int{11, 11, 22, 33, 33, 44, 66}},
	}
	for _, test := range tests {
		t.Run(fmt.Sprint(test.list), func(t *testing.T) {
			exp := slices.Clone(test.list)
			sort.Ints(exp)

			sortFn(test.list)

			require.Equal(t, exp, test.list)
		})
	}
}
