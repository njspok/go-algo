package intersect

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		a   []int
		b   []int
		exp []int
	}{
		{a: []int{}, b: []int{}, exp: []int{}},
		{a: []int{11}, b: []int{}, exp: []int{}},
		{a: []int{11}, b: []int{11}, exp: []int{11}},
		{a: []int{11, 22, 33}, b: []int{11, 22, 33}, exp: []int{11, 22, 33}},
		{a: []int{11, 22, 33, 44}, b: []int{11, 22, 44}, exp: []int{11, 22, 44}},
		{a: []int{11, 22, 33, 11, 22}, b: []int{11, 22, 44, 11}, exp: []int{11, 22, 11}},
	}
	for n, tt := range tests {
		// change args should not change result

		t.Run(fmt.Sprintf("direct %d", n), func(t *testing.T) {
			require.Equal(t, tt.exp, Intersect(tt.a, tt.b))
		})
		t.Run(fmt.Sprintf("reverse %d", n), func(t *testing.T) {
			require.Equal(t, tt.exp, Intersect(tt.b, tt.a))
		})
	}
}
