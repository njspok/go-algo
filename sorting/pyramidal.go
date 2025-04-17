package sorting

import "container/heap"

func PyramidalSort(a []int) {
	h := IntsHeap(a)
	heap.Init(&h)
	for h.Len() > 0 {
		heap.Pop(&h)
	}
}

type IntsHeap []int

func (a *IntsHeap) Len() int { return len(*a) }

func (a *IntsHeap) Less(i, j int) bool {
	return (*a)[i] > (*a)[j]
}

func (a *IntsHeap) Swap(i, j int) {
	(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
}

func (a *IntsHeap) Push(x any) {
	*a = append(*a, x.(int))
}

func (a *IntsHeap) Pop() any {
	size := a.Len()
	last := (*a)[size-1]
	*a = (*a)[:size-1]
	return last
}
