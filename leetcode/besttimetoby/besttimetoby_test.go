package besttimetoby

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func maxProfit(prices []int) int {
	return 0
}

func Test(t *testing.T) {
	tests := []struct {
		prices []int
		result int
	}{
		{nil, 0},
		{[]int{}, 0},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%+v", test.prices), func(t *testing.T) {
			require.Equal(t, test.result, maxProfit(test.prices))
		})
	}
}
