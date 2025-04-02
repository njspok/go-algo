package mergechans

import (
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func TestMergeChans(t *testing.T) {
	ch01 := make(chan int, 3)
	ch01 <- 1
	ch01 <- 2
	ch01 <- 3
	close(ch01)

	ch02 := make(chan int, 3)
	ch02 <- 11
	ch02 <- 22
	ch02 <- 33
	close(ch02)

	ch03 := make(chan int, 3)
	ch03 <- 111
	ch03 <- 222
	ch03 <- 333
	close(ch03)

	var list []int
	for v := range MergeChans(ch01, ch02, ch03) {
		list = append(list, v)
	}

	sort.Sort(sort.IntSlice(list))

	require.Len(t, list, 9)
	require.Equal(t, []int{1, 2, 3, 11, 22, 33, 111, 222, 333}, list)
}
