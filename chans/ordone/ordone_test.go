package ordone

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOrDone(t *testing.T) {
	t.Run("done", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		ch := make(chan int)

		res := OrDone(ctx, ch)

		var list []int
		for n := range res {
			list = append(list, n)
		}
		require.Empty(t, list)
	})
	t.Run("close src chan", func(t *testing.T) {
		ch := make(chan int, 3)
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)

		res := OrDone(context.Background(), ch)

		var list []int
		for n := range res {
			list = append(list, n)
		}
		require.Equal(t, []int{1, 2, 3}, list)
	})
}

func TestWaitReadWriteInChans(t *testing.T) {
	t.Run("read done successfully", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		ch := make(chan int)

		var res string
		select {
		case <-ctx.Done():
			res = "done"
		case ch <- 111:
			res = "111"
		}

		require.Equal(t, "done", res)
	})
	t.Run("write to chan successfully", func(t *testing.T) {
		ctx := context.Background()

		ch := make(chan int, 1)

		var res string
		select {
		case <-ctx.Done():
			res = "done"
		case ch <- 111:
			res = "111"
		}

		require.Equal(t, "111", res)
	})
}
