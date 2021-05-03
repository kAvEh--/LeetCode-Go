package main

func MaxProfit(prices []int) int {
	profit := 0
	low := prices[0]
	i := 0
	for i < len(prices)-1 {
		for i < len(prices)-1 && prices[i] >= prices[i+1] {
			i++
		}
		low = prices[i]
		for i < len(prices)-1 && prices[i] < prices[i+1] {
			i++
		}
		profit += prices[i] - low
	}
	return profit
}