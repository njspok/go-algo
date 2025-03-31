package binsearch

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		require.Equal(t, -1, BinSearch([]int{}, 11))
	})

	t.Run("one", func(t *testing.T) {
		require.Equal(t, 0, BinSearch([]int{11}, 11))
		require.Equal(t, -1, BinSearch([]int{11}, 22))
		require.Equal(t, -1, BinSearch([]int{11}, 5))
	})

	t.Run("two", func(t *testing.T) {
		require.Equal(t, 0, BinSearch([]int{11, 22}, 11))
		require.Equal(t, 1, BinSearch([]int{11, 22}, 22))
		require.Equal(t, -1, BinSearch([]int{11, 22}, 23))
		require.Equal(t, -1, BinSearch([]int{11, 22}, 10))
		require.Equal(t, -1, BinSearch([]int{11, 22}, 15))

	})

	t.Run("many", func(t *testing.T) {
		const size = 50
		list := make([]int, 0, size)
		for i := 0; i < size; i++ {
			list = append(list, i*10)
		}

		t.Run("find all values", func(t *testing.T) {
			for i, v := range list {
				t.Run(fmt.Sprintf("must find %d with index %d", v, i), func(t *testing.T) {
					require.Equal(t, i, BinSearch(list, v))
				})
			}
		})

		t.Run("not find out range", func(t *testing.T) {
			for _, v := range list {
				v := v + 10_000
				t.Run(fmt.Sprintf("must not find %d", v), func(t *testing.T) {
					require.Equal(t, -1, BinSearch(list, v))
				})
			}
		})

		t.Run("not find in range", func(t *testing.T) {
			for _, v := range list {
				v := v/2 + 1
				t.Run(fmt.Sprintf("must not find %d", v), func(t *testing.T) {
					require.Equal(t, -1, BinSearch(list, v))
				})
			}
		})
	})
}
