package can_place_flowers

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
	tests := []struct {
		flowerbed []int
		n         int
		want      bool
	}{
		{[]int{1, 0, 0, 0, 1}, 1, true},
		{[]int{1, 0, 0, 0, 1}, 2, false},
		{[]int{1}, 0, true},
	}
	for _, tt := range tests {
		if got := CanPlaceFlowers(tt.flowerbed, tt.n); got != tt.want {
			t.Errorf("CanPlaceFlowers(%v, %v) = %v, want %v",
				tt.flowerbed, tt.n, got, tt.want)
		}
	}
}
