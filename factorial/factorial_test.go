package factorial

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test(t *testing.T) {
	var tests = []struct {
		n   uint
		res uint
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
		{11, 39916800},
		{12, 479001600},
		{13, 6227020800},
		{14, 87178291200},
		{20, 2432902008176640000},
		{21, 0}, // error overflow
		{22, 0}, // error overflow
	}
	t.Run("recursive", func(t *testing.T) {
		t.Parallel()
		for _, tt := range tests {
			t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
				require.Equal(t, tt.res, RecursiveFactorial(tt.n))
			})
		}
	})
	t.Run("cycle", func(t *testing.T) {
		t.Parallel()
		for _, tt := range tests {
			t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
				require.Equal(t, tt.res, CycleFactorial(tt.n))
			})
		}
	})
}
