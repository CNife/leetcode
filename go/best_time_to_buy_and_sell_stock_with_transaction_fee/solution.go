package best_time_to_buy_and_sell_stock_with_transaction_fee

func MaxProfit(prices []int, fee int) int {
	//n := len(prices)
	//dp := make([][2]int, n)
	//dp[0][0] = -prices[0]
	//for i := 1; i < n; i++ {
	//	dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
	//	dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	//}
	//fmt.Println(dp)
	//return max(dp[n-1][0], dp[n-1][1])

	m, n := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		m, n = max(m, n-prices[i]), max(n, m+prices[i]-fee)
	}
	return max(m, n)
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
