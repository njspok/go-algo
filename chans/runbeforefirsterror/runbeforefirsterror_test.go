package runbeforefirsterror

import (
	"errors"
	"github.com/stretchr/testify/require"
	"sync/atomic"
	"testing"
	"time"
)

func Test(t *testing.T) {
	t.Run("no errors", func(t *testing.T) {
		var counter atomic.Int32

		err := Run([]func() error{
			func() error {
				counter.Add(1)
				return nil
			},
			func() error {
				counter.Add(1)
				return nil
			},
			func() error {
				counter.Add(1)
				return nil
			},
		}...)

		require.NoError(t, err)
		require.Equal(t, int32(3), counter.Load())
	})
	t.Run("with error", func(t *testing.T) {
		var counter atomic.Int32

		err := Run([]func() error{
			func() error {
				counter.Add(1)
				return nil
			},
			func() error {
				counter.Add(1)
				time.Sleep(100 * time.Millisecond)
				return errors.New("first shit happens")
			},
			func() error {
				counter.Add(1)
				time.Sleep(200 * time.Millisecond)
				return errors.New("second shit happens")
			},
			func() error {
				counter.Add(1)
				time.Sleep(300 * time.Millisecond)
				return nil
			},
		}...)

		require.EqualError(t, err, "first shit happens")
		require.Equal(t, int32(4), counter.Load())
	})
}
