// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/

package besttimetoby

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
