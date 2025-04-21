package singleflight

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestSingleFligh(t *testing.T) {
	sf := New()

	result := make(chan int, 4)

	result <- sf.Do("f1", func() int {
		done := make(chan struct{})

		for range 3 {
			go func() {
				<-done
				result <- sf.Do("f1", func() int {
					panic("must not call")
				})
			}()
		}

		close(done)
		time.Sleep(time.Millisecond * 100)
		return 99
	})

	require.Equal(t, 99, <-result)
	require.Equal(t, 99, <-result)
	require.Equal(t, 99, <-result)
	require.Equal(t, 99, <-result)
}

func TestSingle(t *testing.T) {
	s := &single{}

	result := make(chan int, 4)

	result <- s.do(func() int {
		done := make(chan struct{})

		for range 3 {
			go func() {
				<-done
				result <- s.do(func() int {
					panic("must not call")
				})
			}()
		}

		close(done)
		time.Sleep(time.Millisecond * 100)
		return 99
	})

	require.Equal(t, 99, <-result)
	require.Equal(t, 99, <-result)
	require.Equal(t, 99, <-result)
	require.Equal(t, 99, <-result)

}
