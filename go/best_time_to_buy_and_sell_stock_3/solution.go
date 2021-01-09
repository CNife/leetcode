package best_time_to_buy_and_sell_stock_3

//func MaxProfit(prices []int) int {
//	if len(prices) < 2 {
//		return 0
//	}
//
//	dp := make([][5]int, len(prices))
//	dp[0][firstBought] = -prices[0]
//	dp[0][secondBought] = -prices[0]
//
//	for i := 1; i < len(prices); i++ {
//		price := prices[i]
//		dp[i][notYetBought] = dp[i-1][notYetBought]
//		dp[i][firstBought] = max(dp[i-1][notYetBought]-price, dp[i-1][firstBought])
//		dp[i][firstSold] = max(dp[i-1][firstBought]+price, dp[i-1][firstSold])
//		dp[i][secondBought] = max(dp[i-1][firstSold]-price, dp[i-1][secondBought])
//		dp[i][secondSold] = max(dp[i-1][secondBought]+price, dp[i-1][secondSold])
//	}
//
//	return dp[len(prices)-1][secondSold]
//}
//
//const (
//	notYetBought = iota
//	firstBought
//	firstSold
//	secondBought
//	secondSold
//)

func MaxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	var firstSold, secondSold int
	firstBought, secondBought := -prices[0], -prices[0]

	for i := 1; i < n; i++ {
		price := prices[i]
		secondSold = max(secondSold, secondBought+price)
		secondBought = max(secondBought, firstSold-price)
		firstSold = max(firstSold, firstBought+price)
		firstBought = max(firstBought, -price)
	}

	return secondSold
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
