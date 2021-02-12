package lemonade_change

import "testing"

func TestLemonadeChange(t *testing.T) {
	tests := []struct {
		bills []int
		want  bool
	}{
		{[]int{5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5}, true},
		{[]int{5, 5, 5, 10, 20}, true},
		{[]int{5, 5, 10}, true},
		{[]int{10, 10}, false},
		{[]int{5, 5, 10, 10, 20}, false},
	}
	for _, tt := range tests {
		if got := LemonadeChange(tt.bills); got != tt.want {
			t.Errorf("LemonadeChange(%v) = %v, want %v", tt.bills, got, tt.want)
		}
	}
}
