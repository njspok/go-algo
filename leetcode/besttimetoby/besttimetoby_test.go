package besttimetoby

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

// Complexity
// - memory O(1)
// - time O(n)
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices[1:] {
		if price < minPrice {
			minPrice = price
			continue
		}

		profit := price - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}

// Complexity
// - memory O(1)
// - time O(n*n)
func maxProfitTrivial(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	maxProfit := 0
	for i, pbuy := range prices[:len(prices)-1] {
		for _, psel := range prices[i+1:] {
			profit := psel - pbuy
			if maxProfit < profit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}

func Test(t *testing.T) {
	t.Run("fixed", func(t *testing.T) {
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
	})
	t.Run("compared", func(t *testing.T) {
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
