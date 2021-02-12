package candy

import "testing"

func TestCandy(t *testing.T) {
	tests := []struct {
		ratings []int
		want    int
	}{
		{[]int{1, 0, 2}, 5},
		{[]int{1, 2, 2}, 4},
	}
	for _, tt := range tests {
		if got := Candy(tt.ratings); got != tt.want {
			t.Errorf("Candy(%v) = %v, want %v",
				tt.ratings, got, tt.want)
		}
	}
}
