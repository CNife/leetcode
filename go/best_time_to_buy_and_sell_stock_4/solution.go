package best_time_to_buy_and_sell_stock_4

func MaxProfit(k int, prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	dp := make([][]int, len(prices))
	for i := range dp {
		dp[i] = make([]int, 2*k+1)
	}
	for i := 1; i < 2*k; i += 2 {
		dp[0][i] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		for j := 0; j < 2*k-1; j += 2 {
			dp[i][j+1] = max(dp[i-1][j+1], dp[i-1][j]-prices[i])
			dp[i][j+2] = max(dp[i-1][j+2], dp[i-1][j+1]+prices[i])
		}
	}
	return dp[len(prices)-1][2*k]
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
