package arrays_and_strings

func maxProfit(prices []int) int {
	var profit int
	minP := prices[0]

	for _, price := range prices {
		if minP > price {
			minP = price
		}
		if (price - minP) > profit {
			profit = price - minP
		}
	}

	return profit
}
