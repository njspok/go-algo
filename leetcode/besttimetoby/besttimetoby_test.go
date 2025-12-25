package besttimetoby

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

// Complexity
// - memory O(1)
// - time O(n*n)
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	res := 0
	for i, pbuy := range prices[:len(prices)-1] {
		for _, psel := range prices[i+1:] {
			profit := psel - pbuy
			if res < profit {
				res = profit
			}
		}
	}

	return res
}

func Test(t *testing.T) {
	tests := []struct {
		prices []int
		result int
	}{
		{nil, 0},
		{[]int{}, 0},
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 1, 5, 3, 6, 4, 1, 10, 2, 9}, 9},
		{[]int{7, 6, 4, 3, 1}, 0},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%+v", test.prices), func(t *testing.T) {
			require.Equal(t, test.result, maxProfit(test.prices))
		})
	}
}
