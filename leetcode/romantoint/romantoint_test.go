package romantoint

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		s   string
		res int
	}{
		{"", 0},
		{"I", 1},
		{"V", 5},
		{"X", 10},
		{"L", 50},
		{"C", 100},
		{"D", 500},
		{"M", 1000},

		{"IV", 4},
		{"IX", 9},

		{"XL", 40},
		{"XC", 90},

		{"CD", 400},
		{"CM", 900},

		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"XLIII", 43},
		{"LXXXVIII", 88},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			require.Equal(t, tt.res, RomanToInt(tt.s))
		})
	}
}
