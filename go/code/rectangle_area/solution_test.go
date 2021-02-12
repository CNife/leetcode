package rectangle_area

import "testing"

func TestComputeArea(t *testing.T) {
	tests := []struct {
		A, B, C, D, E, F, G, H, want int
	}{
		{-3, 0, 3, 4, 0, -1, 9, 2, 45},
	}
	for _, tt := range tests {
		got := ComputeArea(tt.A, tt.B, tt.C, tt.D, tt.E, tt.F, tt.G, tt.H)
		if got != tt.want {
			t.Errorf("ComputeArea(%v, %v, %v, %v, %v, %v, %v, %v) = %v, want %v",
				tt.A, tt.B, tt.C, tt.D, tt.E, tt.F, tt.G, tt.H, got, tt.want)
		}
	}
}
