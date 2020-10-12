package jump_game_2

import "testing"

func TestJump(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
	}
	for _, tt := range tests {
		if got := Jump(tt.nums); got != tt.want {
			t.Errorf("got Jump(%v) = %v, want %v", tt.nums, got, tt.want)
		}
	}
}
