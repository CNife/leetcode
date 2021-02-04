package sliding_window_median

import "testing"

func TestMedianSlidingWindow(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []float64
	}{
		{
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:    3,
			want: []float64{1, -1, -1, 3, 5, 6},
		},
	}
	for _, tt := range tests {
		got := MedianSlidingWindow(tt.nums, tt.k)
		for i, wantNum := range tt.want {
			gotNum := got[i]
			if wantNum-gotNum >= 1e-5 || gotNum-wantNum >= 1e-5 {
				t.Fatalf("MedianSlidingWindow(%v, %v) = %v, want %v",
					tt.nums, tt.k, got, tt.want)
			}
		}
	}
}
