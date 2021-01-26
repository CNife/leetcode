package number_of_equivalent_domino_pairs

import "testing"

func TestNumEquivDominoPairs(t *testing.T) {
	tests := []struct {
		dominoes [][]int
		want     int
	}{
		{
			dominoes: [][]int{
				{1, 2},
				{2, 1},
				{3, 4},
				{5, 6},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		clonedDominoes := make([][]int, len(tt.dominoes))
		for i := range clonedDominoes {
			clonedDominoes[i] = []int{tt.dominoes[i][0], tt.dominoes[i][1]}
		}

		if got := NumEquivDominoPairs(clonedDominoes); got != tt.want {
			t.Errorf("NumEquivDominoPairs(%v) = %v, want %v",
				tt.dominoes, got, tt.want)
		}
	}
}
