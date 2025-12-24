package firstuniqcharinstring

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		str string
		pos int
	}{
		{"", -1},
		{"a", 0},
		{"aa", -1},
		{"abc", 0},
		{"abca", 1},
		{"abcab", 2},
		{"abcabc", -1},
		{"aabb", -1},
	}
	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			require.Equal(t, test.pos, firstUniqChar(test.str))
		})
	}
}
