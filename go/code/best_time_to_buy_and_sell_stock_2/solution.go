package best_time_to_buy_and_sell_stock_2

func MaxProfit(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}

	min, profitSum, profit := prices[0], 0, 0
	for _, price := range prices[1:] {
		diff := price - min
		if diff > profit {
			profit = diff
		} else {
			profitSum += profit
			profit, min = 0, price
		}
	}

	if profit != 0 {
		profitSum += profit
	}
	return profitSum
}
