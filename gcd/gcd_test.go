package gcd

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		a, b uint
		res  uint
	}{
		{0, 0, 0},
		{1, 1, 1},
		{10, 2, 2},
		{10, 3, 1},
		{640, 400, 80},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d %d", tt.a, tt.b), func(t *testing.T) {
			require.Equal(t, tt.res, RecursiveGCD(tt.a, tt.b))
			require.Equal(t, tt.res, CycleGCD(tt.a, tt.b))
		})
	}
}
