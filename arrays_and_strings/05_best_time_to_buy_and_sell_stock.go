package arrays_and_strings

import "slices"

func maxProfit(prices []int) int {
	var profit int
	var dayProfit int

	for day, price := range prices {
		dayProfit = slices.Max(prices[day:]) - price
		if profit < dayProfit {
			profit = dayProfit
		}
	}

	return profit
}
