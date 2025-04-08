package sorting

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	TestSort(t, SelectionSort)
}

func Test_min(t *testing.T) {
	t.Run("panic", func(t *testing.T) {
		require.Panics(t, func() {
			min(nil)
		})
	})
	t.Run("success", func(t *testing.T) {
		tests := []struct {
			arr []int
			min int
			pos int
		}{
			{arr: []int{1}, min: 1, pos: 0},
			{arr: []int{1, 3, 4, 1, -1, 0, 34}, min: -1, pos: 4},
		}
		for _, tt := range tests {
			t.Run(fmt.Sprint(tt.arr), func(t *testing.T) {
				pos, m := min(tt.arr)
				require.Equal(t, tt.pos, pos)
				require.Equal(t, tt.min, m)
			})
		}
	})
}
