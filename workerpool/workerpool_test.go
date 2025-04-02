package workerpool

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test(t *testing.T) {
	jobs := make(chan int)

	results := NewWorkerPool(
		5,
		jobs,
		func(x int) int { return x * x },
	)

	go func() {
		for i := 0; i < 10; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	actualCalc := make(map[int]int)
	for result := range results {
		actualCalc[result.Value] = result.Result
	}

	require.Equal(t, map[int]int{
		0: 0,
		1: 1,
		2: 4,
		3: 9,
		4: 16,
		5: 25,
		6: 36,
		7: 49,
		8: 64,
		9: 81,
	}, actualCalc)
}
