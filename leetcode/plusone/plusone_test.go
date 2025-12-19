package plusone

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		Input  []int
		Output []int
	}{
		{[]int{}, []int{}},
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{1, 9}, []int{2, 0}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
		{[]int{9}, []int{1, 0}},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v => %v", test.Input, test.Output), func(t *testing.T) {
			require.Equal(t, test.Output, plusOne(test.Input))
		})
	}
}
