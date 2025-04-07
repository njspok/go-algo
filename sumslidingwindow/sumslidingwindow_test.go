package sumslidingwindow

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSumSlidingWindow(t *testing.T) {
	tests := []struct {
		size    uint
		values  []int
		results []int
	}{
		{size: 0, values: []int{1, 2, 3, 4, 5}, results: []int{1, 2, 3, 4, 5}},
		{size: 1, values: []int{1, 2, 3, 4, 5}, results: []int{1, 2, 3, 4, 5}},
		{size: 2, values: []int{1, 2, 3, 4, 5}, results: []int{1, 3, 5, 7, 9}},
		{size: 3, values: []int{1, 2, 3, 4, 5}, results: []int{1, 3, 6, 9, 12}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("size: %d values %v", tt.size, tt.values), func(t *testing.T) {
			msw := NewSumSlidingWindow(tt.size)
			require.Zero(t, msw.Value())

			var res []int
			for _, v := range tt.values {
				msw.Add(v)
				res = append(res, msw.Value())
			}
			require.Equal(t, tt.results, res)
		})
	}

}
