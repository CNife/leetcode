package maximum_points_you_can_obtain_from_cards

import "testing"

func TestMaxScore(t *testing.T) {
	tests := []struct {
		cardPoints []int
		k, want    int
	}{
		{
			cardPoints: []int{1, 2, 3, 4, 5, 6, 1}, k: 3,
			want: 12,
		},
		{
			cardPoints: []int{2, 2, 2}, k: 2,
			want: 4,
		},
		{
			cardPoints: []int{9, 7, 7, 9, 7, 7, 9}, k: 7,
			want: 55,
		},
		{
			cardPoints: []int{1, 1000, 1}, k: 1,
			want: 1,
		},
		{
			cardPoints: []int{1, 79, 80, 1, 1, 1, 200, 1}, k: 3,
			want: 202,
		},
	}
	for _, tt := range tests {
		if got := MaxScore(tt.cardPoints, tt.k); got != tt.want {
			t.Errorf("MaxScore(%v, %v) = %v, want %v",
				tt.cardPoints, tt.k, got, tt.want)
		}
	}
}
