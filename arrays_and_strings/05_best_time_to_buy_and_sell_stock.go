package arrays_and_strings

import (
	"math"
)

func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	maxProfit := 0

	for _, currentPrice := range prices {
		minPrice = min(currentPrice, minPrice)
		maxProfit = max(maxProfit, currentPrice-minPrice)
	}

	return maxProfit
}
