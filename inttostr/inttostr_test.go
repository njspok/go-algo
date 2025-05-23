package inttostr

import (
	"github.com/stretchr/testify/require"
	"math"
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	tests := []int{
		0,
		1, 2, 3, 4, 5, 6, 7, 8, 9,
		11, 12,
		100, 123,
		math.MaxInt,
		-1, -2, -3, -4, -5, -6, -7, -8, -9,
		-11, -12,
		-100, -123,
		math.MinInt,
	}
	for _, n := range tests {
		require.Equal(t, strconv.Itoa(n), IntToSTr(n))
	}
}
