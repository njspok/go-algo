package firstuniqcharinstring

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}

	type rec struct {
		pos   int
		count int
	}

	table := make(map[int32]rec)

	for pos, char := range s {
		if r, ok := table[char]; !ok {
			table[char] = rec{pos: pos, count: 1}
		} else {
			r.count++
			table[char] = r
		}
	}

	res := -1

	for _, r := range table {
		if r.count > 1 {
			continue
		}

		if res == -1 {
			res = r.pos
		} else {
			if res > r.pos {
				res = r.pos
			}
		}
	}

	return res
}

func Test(t *testing.T) {
	tests := []struct {
		str string
		pos int
	}{
		{"", -1},
		{"a", 0},
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
