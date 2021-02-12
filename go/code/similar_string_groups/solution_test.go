package similar_string_groups

import "testing"

func TestNumSimilarGroups(t *testing.T) {
	tests := []struct {
		strs []string
		want int
	}{
		{
			strs: []string{"tars", "rats", "arts", "star"},
			want: 2,
		},
		{
			strs: []string{"omv", "ovm"},
			want: 1,
		},
	}
	for _, tt := range tests {
		if got := NumSimilarGroups(tt.strs); got != tt.want {
			t.Errorf("NumSimilarGroups(%v) = %v, want %v",
				tt.strs, got, tt.want)
		}
	}
}
