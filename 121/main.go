package main

import "fmt"

func maxProfit(prices []int) int {
	maxProfit := 0
	for i := len(prices) - 1; i > 0; i-- {
		sellPrice := prices[i]
		for j := i - 1; j >= 0; j-- {
			buyPrice := prices[j]
			profit := sellPrice - buyPrice
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}

func main() {
	prices := []int{1, 3, 12, 4}
	fmt.Println(maxProfit(prices))
}
