package most_stones_removed_with_same_row_or_column

import "testing"

func TestRemoveStones(t *testing.T) {
	tests := []struct {
		stones [][]int
		want   int
	}{
		{
			stones: [][]int{{0, 1}, {1, 0}, {1, 1}},
			want:   2,
		},
		{
			stones: [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}},
			want:   5,
		},
		{
			stones: [][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}},
			want:   3,
		},
		{
			stones: [][]int{{0, 0}},
			want:   0,
		},
	}
	for _, tt := range tests {
		got := RemoveStones(tt.stones)
		if got != tt.want {
			t.Errorf("RemoveStones(%v) = %v, want %v",
				tt.stones, got, tt.want)
		}
	}
}
