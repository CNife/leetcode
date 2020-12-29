package best_time_to_buy_and_sell_stock_4

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		k      int
		prices []int
		want   int
	}{
		{2, []int{2, 4}, 2},
		{2, []int{3, 2, 6, 5, 0, 3}, 7},
	}
	for _, tt := range tests {
		if got := MaxProfit(tt.k, tt.prices); got != tt.want {
			t.Errorf("MaxProfit(%v, %v) = %v, want %v",
				tt.k, tt.prices, got, tt.want)
		}
	}
}
