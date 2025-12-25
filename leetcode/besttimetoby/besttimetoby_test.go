package besttimetoby

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

func Test(t *testing.T) {
	t.Run("golden", func(t *testing.T) {
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
				require.Equal(t, test.result, maxProfitTrivial(test.prices))
			})
		}
	})
	t.Run("compared with trivial", func(t *testing.T) {
		for range 100 {
			var prices []int
			for range 100 {
				num := rand.Intn(100)
				prices = append(prices, num)
			}

			require.Equal(t, maxProfitTrivial(prices), maxProfit(prices))
		}
	})
}
