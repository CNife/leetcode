package best_time_to_buy_and_sell_stock_3

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{[]int{3, 3, 5, 0, 0, 3, 1, 4}, 6},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1}, 0},
	}
	for _, tt := range tests {
		if got := MaxProfit(tt.prices); got != tt.want {
			t.Errorf("MaxProfit(%v) = %v, want %v",
				tt.prices, got, tt.want)
		}
	}
}
