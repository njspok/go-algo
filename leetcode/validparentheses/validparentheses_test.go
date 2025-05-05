package validparentheses

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		s   string
		res bool
	}{
		{"", false},
		{")", false},
		{"}", false},
		{"]", false},
		{"()", true},
		{"()[]{}", true},
		{"([{}])", true},
		{"())", false},
		{"(]", false},
		//{"aaaa", false},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			require.Equal(t, tt.res, isValid(tt.s))
		})
	}
}
