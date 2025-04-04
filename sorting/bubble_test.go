package sorting

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"slices"
	"testing"
)

func Test(t *testing.T) {
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
			BubbleSort(test.list)
			require.True(t, slices.IsSorted(test.list), test.list)
		})
	}
}
