package palindrome

import (
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		x      int
		result bool
	}{
		{x: 0, result: true},
		{x: 9, result: true},
		{x: 10, result: false},
		{x: 11, result: true},
		{x: 111, result: true},
		{x: 121, result: true},
		{x: 123, result: false},
		{x: 1111, result: true},
		{x: 1311, result: false},
		{x: 1331, result: true},
		{x: 21331, result: false},
		{x: -1, result: false},
		{x: -121, result: false},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.x), func(t *testing.T) {
			require.Equal(t, tt.result, IsPalindrome(tt.x))
		})
	}
}
