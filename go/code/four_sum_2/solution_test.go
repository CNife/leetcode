package four_sum_2

import "testing"

func TestFourSumCount(t *testing.T) {
	tests := []struct {
		a, b, c, d []int
		want       int
	}{
		{
			a:    []int{1, 2},
			b:    []int{-2, -1},
			c:    []int{-1, 2},
			d:    []int{0, 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		if got := FourSumCount(tt.a, tt.b, tt.c, tt.d); got != tt.want {
			t.Errorf("FourSumCount(%v, %v, %v, %v) = %v, want %v",
				tt.a, tt.b, tt.c, tt.d, got, tt.want)
		}
	}
}
