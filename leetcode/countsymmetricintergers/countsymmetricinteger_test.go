package countsymmetricintergers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCountSymmetricIntegers(t *testing.T) {
	tests := []struct {
		low   int
		high  int
		count int
	}{
		{low: 100, high: 1, count: 0},
		{low: 1, high: 100, count: 9},
		{low: 1200, high: 1230, count: 4},
	}
	for _, tt := range tests {
		require.Equal(t, tt.count, CountSymmetricIntegers(tt.low, tt.high))
	}
}

func Test_isSymmetric(t *testing.T) {
	tests := []struct {
		n     int
		isSym bool
	}{
		{n: 0, isSym: false},
		{n: 1, isSym: false},
		{n: 11, isSym: true},
		{n: 99, isSym: true},
		{n: 101, isSym: false},
		{n: 1000, isSym: false},
		{n: 1010, isSym: true},
		{n: 2112, isSym: true},
		{n: 3140, isSym: true},
	}

	for _, tt := range tests {
		require.Equal(t, tt.isSym, isSymmetric(tt.n))
	}
}

func Test_symByDigit(t *testing.T) {
	tests := []struct {
		str string
		res int
	}{
		{str: "", res: 0},
		{str: "0", res: 0},
		{str: "1", res: 1},
		{str: "2", res: 2},
		{str: "3", res: 3},
		{str: "4", res: 4},
		{str: "5", res: 5},
		{str: "6", res: 6},
		{str: "7", res: 7},
		{str: "8", res: 8},
		{str: "9", res: 9},
		{str: "0123", res: 6},
	}

	for _, tt := range tests {
		require.Equal(t, tt.res, sumByDigit(tt.str))
	}
}
