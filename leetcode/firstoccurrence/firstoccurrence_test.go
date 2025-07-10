package firstoccurrence

import (
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		res      int
	}{
		{haystack: "", needle: "", res: -1},
		{haystack: "hello", needle: "", res: -1},
		{haystack: "", needle: "hello", res: -1},

		{haystack: "hello", needle: "fuck", res: -1},
		{haystack: "hel", needle: "fuck", res: -1},

		{haystack: "hello", needle: "l", res: 2},
		{haystack: "hello", needle: "z", res: -1},

		{haystack: "hello", needle: "hel", res: 0},
		{haystack: "hello", needle: "ell", res: 1},
		{haystack: "hello", needle: "llo", res: 2},
		{haystack: "hello", needle: "lo", res: 3},
		{haystack: "hello", needle: "o", res: 4},
		{haystack: "helololollo", needle: "llo", res: 8},
		{haystack: "helololollo", needle: "llod", res: -1},
		{haystack: "121212121212123", needle: "123", res: 12},
	}
	for n, tt := range tests {
		t.Run(strconv.Itoa(n), func(t *testing.T) {
			require.Equal(t, tt.res, strStr(tt.haystack, tt.needle))
		})
	}
}

func Test1(t *testing.T) {
	a := "one"
	b := "one"
	c := "ZoneZ"

	require.True(t, a == b)
	require.True(t, a == c[1:4])
}
