package best_time_to_buy_and_sell_stock_with_transaction_fee

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices    []int
		fee, want int
	}{
		{[]int{1, 3, 2, 8, 4, 9}, 2, 8},
	}
	for _, tt := range tests {
		if got := MaxProfit(tt.prices, tt.fee); got != tt.want {
			t.Errorf("MaxProfit(%v, %v) = %v, want %v",
				tt.prices, tt.fee, got, tt.want)
		}
	}
}
