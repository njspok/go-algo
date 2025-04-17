package sorting

import (
	"container/heap"
	"fmt"
	"github.com/stretchr/testify/require"
	"slices"
	"testing"
)

func PyramidalSort(a []int) []int {
	res := make([]int, 0, len(a))
	h := Ints(a)
	heap.Init(&h)
	for h.Len() > 0 {
		res = append(res, heap.Pop(&h).(int))
	}
	return res
}

type Ints []int

func (a *Ints) Len() int { return len(*a) }

func (a *Ints) Less(i, j int) bool {
	return (*a)[i] < (*a)[j]
}

func (a *Ints) Swap(i, j int) {
	(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
}

func (a *Ints) Push(x any) {
	*a = append(*a, x.(int))
}

func (a *Ints) Pop() any {
	size := a.Len()
	last := (*a)[size-1]
	*a = (*a)[:size-1]
	return last
}

func TestPyramidalSort(t *testing.T) {
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
			require.True(t, slices.IsSorted(PyramidalSort(test.list)))
			fmt.Println(test.list)
		})
	}
}
